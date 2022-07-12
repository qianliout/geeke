package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default()
	// r.GET("/file", FileSrv)
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	reader, err := ZipExcelFile()
	if err != nil {
		fmt.Println("err1 is ", err.Error())
	}
	err = SaveFile(reader, "helloword.zip")

	if err != nil {
		fmt.Println("err2 is ", err.Error())
	}

}

func FileSrv(ctx *gin.Context) {
	path := ctx.Query("path")

	fp, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return
	}
	defer func() {
		if err := fp.Close(); err != nil {
			fmt.Println("err is ", err)
		}
	}()

	stat, err := fp.Stat()
	if err != nil {
		fmt.Println("err is ", err)
		return
	}

	all := stat.Size()

	splits := strings.Split(path, "/")
	filename := path
	if len(splits) > 0 {
		filename = splits[len(splits)-1]
	}
	contentType := "application/octet-stream;application/zip"
	ctx.Header("Content-Disposition", "attachment; filename="+filename) // 指定下载文件名
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Type", contentType)

	extraHeaders := map[string]string{"Content-Disposition": "attachment; filename=" + filename, "Content-Transfer-Encoding": "binary"}

	ctx.DataFromReader(http.StatusOK, all, contentType, fp, extraHeaders)
}

// 把多个excel打包成一个zip文件返回,
func ZipExcelFile() (io.Reader, error) {
	b := new(bytes.Buffer)

	file1 := "/Users/liuqianli/Downloads/WPSOffice.zip"
	file2 := "/Users/liuqianli/Downloads/VirtualBox.zip"
	fd, _ := os.OpenFile(file1, os.O_RDONLY, os.ModePerm)
	defer fd.Close()

	zw := zip.NewWriter(b)

	hdr := zip.FileHeader{Name: "VirtualBox.zip"}
	w, err := zw.CreateHeader(&hdr)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(w, fd)
	if err != nil {
		return nil, err
	}

	hdr2 := zip.FileHeader{Name: "VirtualBox.zip"}
	w2, err := zw.CreateHeader(&hdr2)
	fd2, _ := os.OpenFile(file2, os.O_RDONLY, os.ModePerm)
	defer fd.Close()

	_, err = io.Copy(w2, fd2)
	if err != nil {
		return nil, err
	}

	if err := zw.Flush(); err != nil {
		return nil, err
	}
	if err := zw.Close(); err != nil {
		return nil, err
	}

	return b, nil
}
func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func DeleteFile(filename string) error {
	if checkFileIsExist(filename) {
		return os.Remove(filename)
	}
	return nil
}
func SaveFile(reader io.Reader, filenamePrefix string) error {

	filename := filenamePrefix + ".zip"
	if checkFileIsExist(filename) {
		// 如果文件存在就先删除该文件
		if err := DeleteFile(filename); err != nil {
			return err
		}
	}
	f, err := os.Create(filename) // 创建文件
	if err != nil {
		return err
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	_, err = w.ReadFrom(reader)
	if err != nil {
		return err
	}
	if err := w.Flush(); err != nil {
		return err
	}

	return nil
}
