package merge

import (
	"bufio"
	"fmt"
	"os"

	"outback/leetcode/Nmerge/pipeline"
)

func ExtMerge() {
	const filenameIn = "small.in"
	const fileSize = 512
	const chunkCount = 4
	const filenameOut = "small.out"

	//const filenameIn = "large.in"
	//const fileSize = 80000000
	//const chunkCount = 4
	//const chunkCount = 8
	//const chunkCount = 100
	//const filenameOut = "large.out"

	//createRandomSource(filenameIn, fileSize)

	p := createPipeline(filenameIn, fileSize, chunkCount)
	writeToFile(p, filenameOut)
	printFile(filenameOut)
}

func createPipeline(filename string, fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	pipeline.Init()

	sortResults := make([]<-chan int, 0)

	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)
		sortResults = append(sortResults, pipeline.InMemSort(source))
	}
	return pipeline.MergeN(sortResults...)
}

func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriteSink(writer, p)
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(file, -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count == 100 {
			break
		}
	}
}

func createRandomSource(fineName string, count int) {
	file, err := os.Create(fineName)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	in := pipeline.RandomSource(count)
	pipeline.WriteSink(file, in)
}
