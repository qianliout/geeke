package main

import "fmt"

func main() {
	//a := []int{1, 2, 3}
	a := make([]int, 20)
	copy(a, []int{3, 4, 5})
	fmt.Println(a)
	fmt.Println(len(a))

}
