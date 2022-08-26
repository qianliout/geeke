package main

import (
	"fmt"
)

func main() {
	defaultMap := map[string]string{"hello": "word"}
	a := defaultMap
	a["hello"] = "123"
	fmt.Println(defaultMap["hello"])
	changeMap(a)
	fmt.Println(a["hello"])

	defaultSlice := []int{1, 2, 3}
	b := defaultSlice
	b[0] = 10
	fmt.Println(b[0])

	changSlice(b)
	fmt.Println(b[0])
}

func changSlice(a []int) {
	a[0] = 20
}

func changeMap(a map[string]string) {
	a["hello"] = "changeMap"
}
