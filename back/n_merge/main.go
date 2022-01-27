package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	merge2 "qianliout/leetcode/back/n_merge/merge"
	pipeline2 "qianliout/leetcode/back/n_merge/pipeline"
)

func main() {
	fileIN := "small.in"
	fileOut := "small.out"
	createInfile(fileIN, 1024*1024*8+515)

	// GenerateSortBlock(fileIN, 1024*1024*4, 4, gen)
	// fmt.Println("块文件生成完")

	files := make([]*os.File, 0)
	filenames, _ := GetAllFilename("/Users/liuqianli/Documents/golang/src/outback/leetcode/")
	fmt.Println("文件个数:", len(filenames))
	out := createPipeline(filenames, -1, len(files), 0, merge2.NoSort)
	file, err := os.Create(fileOut)
	if err != nil {
		panic(err)
	}
	pipeline2.WriterSink(file, out)
	printNum(fileOut)

}

func createInfile(filename string, count int) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	in := pipeline2.RandomSource(count)
	writer := bufio.NewWriter(file)
	pipeline2.WriterSink(writer, in)
	writer.Flush()
}

func printNum(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	i := 0

	source := pipeline2.ReaderSource(file, -1)
	for n := range source {
		fmt.Println(n)
		i++
		if i > 20 {
			break
		}
	}
}

// 准备多路归并的管道
func createPipeline(files []string, chunkSize, count, whence int, sortFunc func(<-chan int64) <-chan int64) <-chan int64 {
	outs := make([]<-chan int64, 0)
	for _, file := range files {
		for i := 0; i < count; i++ {
			// 只能在循环中打开文件，这样每个协程拿的就是不同的文件描述符
			file, err := os.Open(file)
			if err != nil {
				fmt.Println("Open file error :  ", err.Error())
				continue
			}
			seek := int64(i*chunkSize) + int64(whence) // 当为-1时表时读全部
			if seek < 0 {
				seek = 0
			}
			if _, err := file.Seek(seek, 0); err != nil {
				fmt.Println("seek error is  ", err.Error(), whence, i*chunkSize)
				continue
			}
			source := pipeline2.ReaderSource(file, chunkSize)
			out := sortFunc(source)
			outs = append(outs, out)
		}
	}
	return merge2.MergeN(outs...)
}

// 准备多路归并的管道
func createPipeline2(filename string, chunkSize, count, whence int, sortFunc func(<-chan int64) <-chan int64) <-chan int64 {
	outs := make([]<-chan int64, 0)
	for i := 0; i < count; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		fmt.Println("whence is ", int64(i*chunkSize)+int64(whence))
		seek := int64(i*chunkSize) + int64(whence)
		if seek < 0 {
			seek = 0
		}
		if _, err := file.Seek(seek, 0); err != nil {
			fmt.Println("seek error is  ", err.Error(), whence, i*chunkSize)
			continue
		}
		source := pipeline2.ReaderSource(file, chunkSize)
		out := sortFunc(source)
		outs = append(outs, out)
	}
	return merge2.MergeN(outs...)
}

// 生成排序块
func GenerateSortBlock(originalFile string, chunkSize, workerCount int, gentNewFileFunc func() string) {
	file, err := os.Open(originalFile)
	if err != nil {
		panic(err)
	}
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	seeked, length := 0, stat.Size()

	for seeked < int(length) {
		out := createPipeline2(originalFile, chunkSize, workerCount, seeked, merge2.MemorySort)
		blockFilename := gentNewFileFunc()
		blockFile, err := os.Create(blockFilename)
		if err != nil {
			panic(err)
		}
		writer := bufio.NewWriter(blockFile)
		pipeline2.WriterSink(writer, out)
		if err := writer.Flush(); err != nil {
			fmt.Println("GenerateSortBlock", err.Error())
		}
		seeked += chunkSize * workerCount
		fmt.Println("生成了一个,seeked:", seeked)
	}
}

// 获取所有排序块的文件名
func GetAllFilename(pathname string) ([]string, error) {
	res := make([]string, 0)
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		return nil, err
	}
	for _, fi := range rd {
		if strings.Contains(fi.Name(), "block") {
			res = append(res, fi.Name())
			fmt.Println(fi.Name())
		}
	}
	return res, nil
}

func genName() string {
	return fmt.Sprintf("block-%d", time.Now().UnixNano())
}
