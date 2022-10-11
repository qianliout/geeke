package main

import (
	"fmt"
)

func main() {
	s := "abacb"
	res := removeDuplicateLetters(s)
	fmt.Println("res is ", res)
}

func removeDuplicateLetters(s string) string {
	ss := []byte(s)
	lastIndex := make(map[byte]int)
	for i := range ss {
		lastIndex[ss[i]] = i
	}
	stark := make([]byte, 0)
	visit := make(map[byte]bool)
	for i := range ss {
		if visit[ss[i]] {
			continue
		}

		for len(stark) > 0 && stark[len(stark)-1] > ss[i] && lastIndex[stark[len(stark)-1]] > i {
			top := stark[len(stark)-1]
			stark = stark[:len(stark)-1]
			visit[top] = false
		}
		stark = append(stark, ss[i])
		visit[ss[i]] = true
	}
	return string(stark)
}
