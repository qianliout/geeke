package main

import (
	"fmt"
)

func main() {
	secret := "1122"
	guess := "1222"
	hint := getHint(secret, guess)
	fmt.Println("hint ", hint)
}

func getHint(secret string, guess string) string {
	se := make(map[byte]int)
	gu := make(map[byte]int)
	for i := range secret {
		se[secret[i]]++
	}
	for i := range guess {
		gu[guess[i]]++
	}
	a, b := 0, 0
	// 先计算好相同的
	for i := range guess {
		if i < len(secret) && guess[i] == secret[i] {
			a++
			se[guess[i]]--
		}
	}
	// 再计算位置不同的,这两步一定要分开
	for i := range guess {
		if i < len(secret) && guess[i] == secret[i] {
			continue
		}
		if v, ok := se[guess[i]]; ok && v > 0 {
			se[guess[i]]--
			b++
		}
	}

	return fmt.Sprintf("%dA%dB", a, b)
}
