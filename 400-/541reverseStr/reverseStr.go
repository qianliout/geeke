package main

import (
	"fmt"
)

func main() {
	s := "abcdefg"
	str := reverseStr(s, 8)
	fmt.Println(str)
}

func reverseStr2(s string, k int) string {
	ss := []byte(s)
	use := 0
	for i := 0; i < len(ss); i++ {
		use++
		if use == 2*k {
			reverse(ss[i-2*k+1 : i-k+1])
			use = 0
		}
	}

	if use >= k && use < 2*k {
		reverse(ss[len(ss)-1-use+1 : len(ss)-1-use+1+k])
	}
	if use < k {
		reverse(ss[len(ss)-use:])
	}

	return string(ss)
}

func reverse(s []byte) {
	start, end := 0, len(s)-1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start, end = start+1, end-1
	}
}

func reverseStr(s string, k int) string {
	t := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		sub := t[i:min(i+k, len(s))]
		for j, n := 0, len(sub); j < n/2; j++ {
			sub[j], sub[n-1-j] = sub[n-1-j], sub[j]
		}
	}
	return string(t)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
