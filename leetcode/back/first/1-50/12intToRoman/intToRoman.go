package main

import (
	"fmt"
)

func main() {
	fmt.Println(intToRoman(58))
}

func intToRoman(num int) string {
	if num < 0 {
		return ""
	}
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	romanNum := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	start := 0
	ans := make([]string, 0)

	for num > 0 {
		for num >= romanNum[start] {
			num -= romanNum[start]
			ans = append(ans, roman[start])
		}
		start++
	}
	ansStr := ""
	for _, s := range ans {
		ansStr = ansStr + s
	}

	return ansStr

}
