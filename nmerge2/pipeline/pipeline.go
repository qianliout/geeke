package pipeline

import (
	"encoding/binary"
	"fmt"
	"io"
	"time"
)

var startTime time.Time

func Init() {
	startTime = time.Now()
}

func ReaderSource(reader io.Reader, chunkSize int) <-chan int {
	out := make(chan int, 0)
	go func() {
		defer close(out)
		size := 0
		buffer := make([]byte, 8)
		for {
			n, err := reader.Read(buffer)
			if n > 0 {
				size += n
				fmt.Println("size is ", size, chunkSize)
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			if err != nil || (chunkSize != -1 && size >= chunkSize) {
				break
			}
		}
	}()

	return out
}

func WriterSink(writer io.Writer, in <-chan int) {
	for n := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(n))
		writer.Write(buffer)
	}
}
