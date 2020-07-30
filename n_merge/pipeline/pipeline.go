package pipeline

import (
	"encoding/binary"
	"io"
	"math/rand"
)

func ReaderSource(reader io.Reader, chunkSize int) <-chan int64 {
	out := make(chan int64, 0)
	go func() {
		defer close(out)
		readSize := 0
		for {
			buf := make([]byte, 8)
			n, err := reader.Read(buf)
			if n > 0 {
				readSize += n
				v := binary.BigEndian.Uint64(buf)
				out <- int64(v)
			}
			if err != nil || (chunkSize != -1 && readSize >= chunkSize) {
				break
			}
		}
	}()

	return out

}

func WriterSink(writer io.Writer, in <-chan int64) {
	for n := range in {
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(n))
		writer.Write(buf)
	}
}

func RandomSource(count int) <-chan int64 {
	out := make(chan int64, 0)

	go func() {
		defer close(out)
		for i := 0; i < count; i++ {
			n := rand.Int63()
			out <- n
		}
	}()
	return out
}
