package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"outback/leetcode/n_merge/merge"
	"outback/leetcode/n_merge/pipeline"
)

func main() {
	fileIN := "small.in"
	fileOut := "small.out"
	createInfile(fileIN, 1024*1024*80)
	gen := func() string {
		return fmt.Sprintf("block-%d", time.Now().UnixNano())
	}
	GenerateSortBlock(fileIN, 1024*256, 32, gen)
	fmt.Println("块文件生成完")
	fiels, _ := GetAllFilename("/Users/liuqianli/Documents/golang/src/outback/leetcode/")
	out := createPipeline(fiels, -1, len(fiels), 0, func(in <-chan int64) <-chan int64 {
		return in
	})
	file, err := os.Create(fileOut)
	if err != nil {
		panic(err)
	}

	pipeline.WriterSink(file, out)

	printNum(fileOut)
}

func createInfile(filename string, count int) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	in := pipeline.RandomSource(count)
	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, in)
	writer.Flush()
}

func printNum(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	i := 0

	source := pipeline.ReaderSource(file, -1)
	for n := range source {
		fmt.Println(n)
		i++
		if i > 20 {
			break
		}
	}
}

func createPipeline(files []string, chunkSize, count, whence int, sortFunc func(<-chan int64) <-chan int64) <-chan int64 {

	outs := make([]<-chan int64, 0)
	for _, f := range files {
		file, err := os.Open(f)
		if err != nil {
			panic(err)
		}
		for i := 0; i < count; i++ {
			fmt.Println("whence is ", int64(i*chunkSize)+int64(whence))
			if _, err := file.Seek(int64(i*chunkSize)+int64(whence), 0); err != nil {
				fmt.Println("seek error is  ", err.Error(), whence, i*chunkSize)
				continue
			}
			source := pipeline.ReaderSource(file, chunkSize)
			out := sortFunc(source)
			outs = append(outs, out)
		}
	}
	return merge.MergeN(outs...)
}

// 生成排序块
func GenerateSortBlock(originalFile string, singleSize, workerCount int, gentNewFileFunc func() string) {
	file, err := os.Open(originalFile)
	if err != nil {
		panic(err)
	}
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	seeked, length := 0, stat.Size()
	fmt.Println("length is ", length)

	for seeked < int(length) {
		out := createPipeline2(file, singleSize, workerCount, seeked, merge.MemorySort)
		nf := gentNewFileFunc()
		// GenerateSortBlock2(out, nf)
		file, err := os.Create(nf)
		if err != nil {
			panic(err)
		}
		writer := bufio.NewWriter(file)
		pipeline.WriterSink(writer, out)
		if err := writer.Flush(); err != nil {
			fmt.Println("GenerateSortBlock", err.Error())
		}
		seeked += singleSize * workerCount
	}
}

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
