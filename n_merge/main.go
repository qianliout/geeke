package main

import (
	"fmt"
	"os"

	"outback/leetcode/n_merge/merge"
	"outback/leetcode/n_merge/pipeline"
)

func main() {
	fileIN := "small.in"
	fileOut := "small.out"
	createInfile(fileIN, 64)
	out := createPipeline(fileIN, 512, 4)
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
	pipeline.WriterSink(file, in)
}

func createPipeline(filename string, size, count int) <-chan int64 {
	chunkSize := size / count
	outs := make([]<-chan int64, 0)
	for i := 0; i < count; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReaderSource(file, chunkSize)
		out := merge.InmSort(source)
		outs = append(outs, out)
	}
	return merge.MergeN(outs...)
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
		if i > 10 {
			break
		}
	}
}
