package main

import (
	"fmt"
	"time"
	"unsafe"
)

type Part1 struct {
	a bool
	b int32
	c int8
	d int64
	e byte
	f string
}

const name = "helsl"

func IsNilOrNot(v interface{}) bool {
	fmt.Printf("v size %d\n ", unsafe.Sizeof(v))
	return v == nil
}

func main() {
	go func() {
		i := 0
		for {
			i++
		}
	}()

	time.Sleep(time.Second)
	fmt.Println("fuck")
}
