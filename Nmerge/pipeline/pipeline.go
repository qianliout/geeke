package pipeline

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

var startTime time.Time

func Init() {
	startTime = time.Now()
}

func InMemSort(in <-chan int) <-chan int {
	// add buffer
	out := make(chan int, 1024)
	go func() {
		// Read
		a := make([]int, 0)
		for v := range in {
			a = append(a, v)
		}
		fmt.Println("Read done:", time.Now().Sub(startTime))

		// Sort
		sort.Ints(a)
		fmt.Println("Sort done:", time.Now().Sub(startTime))

		// Print
		for _, v := range a {
			out <- v
		}

		close(out)
	}()
	return out
}

func RandomSource(count int) <-chan int {
	out := make(chan int)
	go func(count int) {
		for i := 0; i < count; i++ {
			n := rand.Int()
			out <- n
		}
		close(out)
	}(count)
	return out
}

func Merge(ch1, ch2 <-chan int) chan int {
	out := make(chan int)
	go func() {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		for ok1 || ok2 {
			if ok1 || (ok2 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-ch1
			} else {
				out <- v2
				v2, ok2 = <-ch2
			}
		}
		close(out)
	}()
	return out
}

func MergeN(inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}
	m := len(inputs) / 2
	return Merge(MergeN(inputs[:m]...), MergeN(inputs[m:]...))
}

// 写数据，本机上一个int占8位
func WriteSink(writer io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

// chunkCount表示最大读取的字节数
func ReaderSource(reader io.Reader, chunkCount int) <-chan int {
	out := make(chan int)
	go func() {
		bytesRead := 0
		for {
			buffer := make([]byte, 8)
			n, err := reader.Read(buffer)
			if err != nil || (chunkCount != -1 && bytesRead >= chunkCount) {
				break
			}
			if n > 0 {
				bytesRead += n
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
		}
		close(out)
	}()
	return out
}
