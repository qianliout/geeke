package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/minio/minio-go"
)

type shell struct {
	S []string
}

func main() {
	// // 源档案（准备压缩的文件或目录）
	// var src = "log"
	// // 目标文件，压缩后的文件
	// var dst = "log.zip"
	//
	// buf := Tar(dst, src)
	// saveToS3(dst, dst, buf)

	// downChan := make(chan bool)
	// defer close(downChan)
	// numChan := genNum(10)
	// donw(numChan, downChan)
	// select {
	// case <-downChan:
	// }
	// fmt.Println("down")
	// s := new(shell)
	// s.S = append(s.S, "fuck")
	// fmt.Println(s.S)
	//
	// m := make(map[int]int)
	// m[3] = 6
	// fmt.Printf("%p\n", m)
	// change(m)
	// fmt.Println(m[3])
	// fmt.Printf("%p", m)
	// var m map[string]int
	// var s []int
	// fmt.Println(len(m))
	// fmt.Println(len(s))
	// s = append(s,3)
	// fmt.Println(len(s))
	// var i byte
	// go func() {
	// 	for i = 0; i <= 255; i++ {
	// 		// fmt.Println("i si ",i)
	// 	}
	// }()
	// fmt.Println("Dropping mic")
	// // Yield execution to force executing other goroutines
	// runtime.Gosched()
	// runtime.GC()
	// fmt.Println("Done")

	time.Parse("2020-10-28T09:02:44")

}

func genNum(num int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; i < num; i++ {
			out <- i
			time.Sleep(time.Second)
		}
	}()
	return out
}

func change(m map[int]int) {
	for i := 0; i < 30000; i++ {
		m[i] = i
	}
}

func donw(numChan chan int, downChan chan bool) {

	go func() {
		defer func() {
			downChan <- true
			println("close down chan")
		}()
		for {
			num, open := <-numChan
			if !open {
				break
			}
			fmt.Println("num is ", num)
		}
	}()
	fmt.Println("返回了")
}

var HeaderTitlesHead = []string{
	"对话ID",
	"访客ID",
	"有效对话",
	"遗漏对话",
	"对话分级",
	"小结",
	"客服姓名",
	"小结时间",
	"地理位置",
}

func Tar(dst, src string) *bytes.Buffer {
	var b bytes.Buffer

	zw := zip.NewWriter(&b)

	defer func() {
		if err := zw.Flush(); err != nil {
			log.Println(err)
		}

		if err := zw.Close(); err != nil {
			log.Println(err)
		}
	}()

	for i := 10; i < 15; i++ {
		ex := genFile(strconv.Itoa(i))
		hdr := zip.FileHeader{Name: ex.Path}
		w, err := zw.CreateHeader(&hdr)
		if err != nil {
			log.Println("WriteHeader", err.Error())
		}

		b, err := ex.WriteToBuffer()
		if err != nil {
			log.Println("WriteToBuffer", err.Error())
		}
		n, err := io.Copy(w, b)

		// n, err := tw.Write(b.Bytes())
		// n, err := fw.Write(b.Bytes())
		if err != nil {
			log.Println("io.Copy", err.Error())
		}
		log.Println("copy n ", n)
	}
	return &b
}

func saveToS3(dst, objName string, buf *bytes.Buffer) error {
	// file, err := os.Open(dst)
	// defer file.Close()
	// if err != nil {
	//     return err
	// }
	//
	// bys, err := ioutil.ReadAll(file)
	// if err != nil {
	//     return err
	// }
	//
	// buf := bytes.NewBuffer(bys)
	opts := minio.PutObjectOptions{ContentType: "application/octet-stream"}
	client, _ := NewClient()
	if _, err := client.PutObject("export", objName, buf, int64(buf.Len()), opts); err != nil {
		return err
	}
	return nil
}

func genFile(fileName string) *excelize.File {
	sheetName := "helloword"
	// write excel header
	excelFile := excelize.NewFile()
	excelFile.SetSheetName("Sheet1", sheetName)
	// 文件名
	excelFile.Path = fileName + ".xlsx"
	err := excelFile.SetSheetRow(sheetName, "A1", &HeaderTitlesHead)
	// 写入表头
	if err != nil {
		log.Println("SetSheetRow", err.Error())
		return nil
	}
	return excelFile
}

func NewClient() (*minio.Client, error) {

	surl := "http://10.52.4.40:9010"
	region := ""
	access_key_id := "admin"
	secret_access_key := "123123123"

	u, err := url.Parse(surl)
	if err != nil {
		log.Println("client err ", err)
	}

	client, err := minio.NewWithRegion(u.Host, access_key_id, secret_access_key, u.Scheme == "https", region)
	if err != nil {
		log.Println("client err ", err)
	}
	return client, err
}
