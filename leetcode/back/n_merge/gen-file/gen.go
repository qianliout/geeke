package main

import (
	"bufio"
	"encoding/binary"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Create("/home/liuqiang/tmp/bigdata/bigdata.in")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	b := bufio.NewWriter(file)
	defer b.Flush()

	n := 1000000000 // 8G
	WriterSink(b, Randomw(n))
}

func Randomw(count int) <-chan int64 {
	out := make(chan int64)
	go func() {
		defer close(out)
		for i := 0; i < count; i++ {
			out <- time.Now().UnixNano()
		}
	}()
	return out
}

func WriterSink(writer io.Writer, in <-chan int64) {
	for n := range in {
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(n))
		if _, err := writer.Write(buf); err != nil {
			log.Println(err.Error())
		}
	}
}
