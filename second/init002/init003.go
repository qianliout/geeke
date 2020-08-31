package main

import (
	"fmt"
)

func main() {
	fmt.Println(len([]rune("脸在那里")))
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
