package main

import (
	"fmt"
)

func main() {
	s := "AAAAAAAAAAA"
	sequences := findRepeatedDnaSequences(s)
	fmt.Println(sequences)
}

func findRepeatedDnaSequences(s string) []string {
	ans, ss, exit := make([]string, 0), []byte(s), make(map[string]int)

	for i := 0; i <= len(ss)-10; i++ {
		exit[string(ss[i:i+10])] += 1
	}
	for k, v := range exit {
		if v >= 2 {
			ans = append(ans, k)
		}
	}
	return ans
}
