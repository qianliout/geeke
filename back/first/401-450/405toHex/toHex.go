package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	res := toHex2(-1)
	fmt.Println("res is ", res)
}

/*
给定一个整数，编写一个算法将这个数转换为十六进制数。对于负整数，我们通常使用 补码运算 方法。
注意:
    十六进制中所有字母(a-f)都必须是小写。
    十六进制字符串中不能包含多余的前导零。如果要转化的数为0，那么以单个字符'0'来表示；对于其他情况，十六进制字符串中的第一个字符将不会是0字符。
    给定的数确保在32位有符号整数范围内。
    不能使用任何由库提供的将数字直接转换或格式化为十六进制的方法。
示例 1：
输入:
26
输出:
"1a"
示例 2：
输入:
-1
输出:
"ffffffff"
*/

/*
1、当正数时，直接循环右移4位（除以16也可以），直到商为0，再依次将余数字符串起来，逆序排列。
2、当0时，直接返回0。
3、当负数时，将其绝对值通过与（2的32次方减1）异或再加1，转化为对应的十进制正数，然后执第一步的操作。
注：设定一个字符列表来替换while中的if elif判断也挺不错。
*/

// 这种解法是错的，我也不知道错在什么地方
func toHex(num int) string {
	if num == 0 {
		return ""
	}
	if num < 0 {
		num = int(math.Abs(float64(num))) ^ int(math.Pow(float64(2), float64(32))) + 1
	}
	res := make([]string, 0)
	for num>>4 > 0 || num > 0 {
		last := strconv.Itoa(num % 16)
		switch last {
		case "10":
			last = "a"
		case "11":
			last = "b"
		case "12":
			last = "c"
		case "13":
			last = "d"
		case "14":
			last = "e"
		case "15":
			last = "f"
		}
		num = num >> 4
		res = append(res, last)
	}
	ans := ""
	for i := len(res) - 1; i >= 0; i-- {
		ans = ans + res[i]
	}
	return ans
}

// 只适用于32位，
func toHex2(num int) string {
	if num < 0 {
		num += 4294967296
	}
	if num == 0 {
		return "0"
	}
	res := []int32{}
	hash := []int32{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	for num != 0 {
		temp := num % 16
		num = num / 16
		res = append([]int32{hash[temp]}, res...)
	}
	return string(res)
}
