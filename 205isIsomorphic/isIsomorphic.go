package main

import (
	"fmt"
)

func main() {
	s := isIsomorphic2("egg", "add")
	// s := isIsomorphic2("paper", "title") // true 说明只有 s到t就行
	fmt.Println(s)
}

func isIsomorphic(s string, t string) bool {
	return compare(s, t) && compare(t, s)
}

func isIsomorphic2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash1, hash2 := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s); i++ {
		hash1[s[i]] = i
		hash2[t[i]] = i
	}
	for i := 0; i < len(s); i++ {
		if hash1[s[i]] != hash2[t[i]] {
			return false
		}
	}
	return true
}

// 从s到t
func compare(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	map1 := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if v, ok := map1[s[i]]; ok {
			map1[s[i]] = t[i]
		} else {
			if v != t[i] {
				return false
			}
		}
	}
	return true
}
