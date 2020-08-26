package main

import (
	"fmt"
)

func main() {
	 h.initF()
}

type Hello struct {
	initF func() int
}

var h Hello


func SayHello() int {
	fmt.Println("heello wordc")
	return 3
}

func init() {
	initF := func() int {
		return SayHello()
	}
	fmt.Println("3333")
	h = Hello{initF: initF}
}
