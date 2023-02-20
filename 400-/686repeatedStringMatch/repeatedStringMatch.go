package main

import (
	"strings"
)

func main() {

}

func repeatedStringMatch(a string, b string) int {
	ti := len(b) * len(a)
	for i := 0; i < 3; i++ {
		if strings.Contains(strings.Repeat(a, ti+i), b) {
			return ti + 1
		}
	}
	return -1
}
