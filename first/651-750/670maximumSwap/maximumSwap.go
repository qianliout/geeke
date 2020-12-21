package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(maximumSwap(98368))
	// fmt.Println(maximumSwap(9973))
	// fmt.Println(maximumSwap(1234))
}

/*
给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。
示例 1 :
输入: 2736
输出: 7236
解释: 交换数字2和数字7。
示例 2 :
输入: 9973
输出: 9973
解释: 不需要交换。
*/
func maximumSwap(num int) int {
	ss := []byte(strconv.Itoa(num))
	ss2 := append([]byte{}, ss...)
	sort.Slice(ss2, func(i, j int) bool { return ss2[i] > ss2[j] })
	for i := 0; i < len(ss); i++ {
		if ss[i] != ss2[i] {
			for j := len(ss) - 1; j >= i; j-- {
				if ss[j] == ss2[i] {
					ss[i], ss[j] = ss[j], ss[i]
					atoi, _ := strconv.Atoi(string(ss))
					return atoi
				}
			}
		}
	}
	return num

}
