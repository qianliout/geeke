package main

import (
	"fmt"
)

func main() {
	s := "aabcccccaaabcaaa"
	res := compressString2(s)
	fmt.Println("ans is ", res)
}

/*
 字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2b1c5a3。
若“压缩”后的字符串没有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。
示例1:
 输入："aabcccccaaa"
 输出："a2b1c5a3"
示例2:
 输入："abbccd"
 输出："abbccd"
 解释："abbccd"压缩后为"a1b2c2d1"，比原字符串长度更长。
提示：
    字符串长度在[0, 50000]范围内。
*/
func compressString(S string) string {
	ans := ""
	if len(S) <= 0 {
		return ans
	}
	byteSlice := []byte(S)
	var pre byte = byteSlice[0]
	num := 0
	for i, b := range byteSlice {
		if b == pre {
			num++
		} else if b != pre {
			ans = fmt.Sprintf("%s%s%d", ans, string(pre), num)
			num = 1
			pre = b
		}
		if i == len(byteSlice)-1 {
			ans = fmt.Sprintf("%s%s%d", ans, string(pre), num)
		}
	}
	if len(ans) < len(S) {
		return ans
	}
	return S
}

func compressString2(S string) string {
	ans := ""
	if len(S) <= 0 {
		return ans
	}
	sliceS := []byte(S)
	pre := sliceS[0]
	left, right := 0, 0
	for right < len(sliceS) {
		if sliceS[right] == pre {
			right++
		} else {
			ans = fmt.Sprintf("%s%s%d", ans, string(pre), right-left)
			left = right
			pre = sliceS[right]
			right++
		}

		if right >= len(sliceS) {
			ans = fmt.Sprintf("%s%s%d", ans, string(pre), right-left)
		}
	}
	if len(ans) < len(S) {
		return ans
	}
	return S
}
