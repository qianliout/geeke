package main

import (
	"fmt"
	"math"
)

func main() {
	s := countAndSay(4)
	fmt.Println("sya is ", s)
}

/*
给定一个正整数 n ，输出外观数列的第 n 项。
「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。
你可以将其视作是由递归公式定义的数字字符串序列：
countAndSay(1) = "1"
countAndSay(n) 是对 countAndSay(n-1) 的描述，然后转换成另一个数字字符串。
*/
func countAndSay2(n int) string {
	ans := "1"
	for i := 2; i <= n; i++ {
		ans = say(ans)
	}
	return ans
}

func countAndSay(n int) string {
	ints := sayint(n)
	ans := ""
	for _, b := range ints {
		ans = fmt.Sprintf("%s%d", ans, b)
	}
	return ans
}

func say(s string) string {
	s = fmt.Sprintf("%s+", s)
	ss := []byte(s)
	start := 1
	ans := ""
	for i := 0; i < len(ss); i++ {
		if i > 0 && ss[i] == ss[i-1] {
			start++
		} else if i > 0 && ss[i] != ss[i-1] {
			ans = fmt.Sprintf("%s%d%d", ans, start, ss[i-1]-'0')
			start = 1
		}
	}

	return ans
}

func sayint(n int) []int {
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{1}
	}
	ss := sayint(n - 1)
	ss = append(ss, math.MinInt64)

	start := 1
	ans := make([]int, 0)
	for i := 0; i < len(ss); i++ {
		if i > 0 && ss[i] == ss[i-1] {
			start++
		} else if i > 0 && ss[i] != ss[i-1] {
			ans = append(ans, start, ss[i-1])
			start = 1
		}
	}
	return ans
}
