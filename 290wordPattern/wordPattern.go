package main

import (
	"fmt"
	"strings"
)

func main() {
	pattern := "abba"
	s := "dog cat cat dog"
	// s := "dog dog dog dog"
	b := wordPattern(pattern, s)
	fmt.Println(b)
}

func wordPattern2(pattern string, s string) bool {
	pp, ss := []byte(pattern), strings.Split(s, " ")
	exit1 := make(map[byte]string)
	exit2 := make(map[string]byte)
	if len(pp) != len(ss) {
		return false
	}
	for i := range pp {
		k1, ok1 := exit1[pp[i]]
		k2, ok2 := exit2[ss[i]]
		if !ok1 && !ok2 {
			exit1[pp[i]] = ss[i]
			exit2[ss[i]] = pp[i]
		} else if ok1 && ok2 && k1 == ss[i] && k2 == pp[i] {
			continue
		} else {
			return false
		}
	}
	return true
}

func wordPattern(pattern string, s string) bool {
	pp, ss := []byte(pattern), strings.Split(s, " ")
	exit1 := make(map[byte]int)
	exit2 := make(map[string]int)
	if len(pp) != len(ss) {
		return false
	}
	for i := range pp {
		v1, ok1 := exit1[pp[i]]
		v2, ok2 := exit2[ss[i]]
		if (!ok1 && !ok2) || (ok1 && ok2 && v1 == v2) {
			exit1[pp[i]] = i
			exit2[ss[i]] = i
		} else {
			return false
		}
	}
	return true
}
