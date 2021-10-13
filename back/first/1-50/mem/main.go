package main

import (
	"context"
	"fmt"
	"strconv"
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
	itoa := strconv.Itoa(123232)
	a := string(132323)
	fmt.Printf("a is %T,%T", a, itoa)
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
}
