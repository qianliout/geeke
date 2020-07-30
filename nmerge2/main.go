package main

import (
	"bufio"
	"os"

	"outback/leetcode/nmerge2/merge"
	"outback/leetcode/nmerge2/pipeline"
	"outback/leetcode/nmerge2/source"
)

func main() {
	merge.Init()
	fileIn := "small.in"
	fileOut := "small.out"
	// 生成数据文件
	createInFile(fileIn, 64)

	out := createPipeline(fileIn, 512, 4)
	// 读取数据文件并排序
	// 排序后的内存写回文件
	writeToFile(fileOut, out)
}

func createInFile(filename string, count int) {
	in := source.RandomSource(count)
	writeToFile(filename, in)
}

func createPipeline(filename string, size, chunk int) <-chan int {

	chunkSize := size / chunk
	sorts := make([]<-chan int, 0)

	for i := 0; i < chunk; i++ {
		file, err := os.Open(filename)
		defer file.Close()
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)
		s := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)
		sorts = append(sorts, merge.InmSort(s))
	}
	out := merge.MergeN(sorts...)
	return out
}

func writeToFile(filename string, in <-chan int) {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriterSink(writer, in)
}
