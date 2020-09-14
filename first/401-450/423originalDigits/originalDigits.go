package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res := originalDigits("owoztneoerfviefuroonetwo")
	fmt.Println("res is ", res)
}

/*
给定一个非空字符串，其中包含字母顺序打乱的英文单词表示的数字0-9。按升序输出原始的数字。
注意:
    输入只包含小写英文字母。
    输入保证合法并可以转换为原始的数字，这意味着像 "abc" 或 "zerone" 的输入是不允许的。
    输入字符串的长度小于 50,000。
示例 1:
输入: "owoztneoer"
输出: "012" (zeroonetwo)
示例 2:
输入: "fviefuro"
输出: "45" (fourfive)
*/

// 保证输入是合法有效的
/*
规律如下：
(1)['zero','two','four','six','eight','one','three','five','seven','nine']，观察这十个单词会发现，
['zero','two','four','six','eight']这五个单词具有唯一字母['z', 'w', 'u', 'x', 'g']，所以可以先筛选出这五个字母根据s的频数统计表可以通过这五个字母的频数计算出所对应的单词数，然后在s的频数统计表中删除这五个单词所有字母的频数；
(2)接下来可以发现在上一步删除完成后，['one','three','five','seven']这四个单词中['o','r', 'f', 's']成为了唯一字母，于是同上一步操作；
(3)最后['nine']中的['i']也成为了唯一字母，继续(1)的操作
*/

func originalDigits(s string) string {
	byteMap := map[byte]string{
		'o': "one",
		'w': "two",
		'r': "three",
		'u': "four",
		'f': "five",
		'x': "six",
		's': "seven",
		'g': "eight",
		'i': "nine",
		'z': "zero",
	}

	numMap := map[byte]int{
		'o': 1,
		'w': 2,
		'r': 3,
		'u': 4,
		'f': 5,
		'x': 6,
		's': 7,
		'g': 8,
		'i': 9,
		'z': 0,
	}
	ansMap := make(map[int]int)
	if len(s) == 0 {
		return ""
	}

	ss := []byte(s)
	exit := make(map[byte]int)
	for _, ch := range ss {
		exit[ch]++
	}
	for _, ch := range []byte{'z', 'w', 'u', 'x', 'g'} {
		zn := exit[ch]
		for _, i := range []byte(byteMap[ch]) {
			exit[i] = exit[i] - zn
		}
		ansMap[numMap[ch]] = zn
	}
	for _, ch := range []byte{'o', 'r', 'f', 's'} {
		zn := exit[ch]
		for _, i := range []byte(byteMap[ch]) {
			exit[i] = exit[i] - zn
		}
		ansMap[numMap[ch]] = zn
	}
	for _, ch := range []byte{'i'} {
		zn := exit[ch]
		for _, i := range []byte(byteMap[ch]) {
			exit[i] = exit[i] - zn
		}
		ansMap[numMap[ch]] = zn
	}
	ans := ""
	for i := 0; i <= 9; i++ {
		ans = ans + strings.Repeat(strconv.Itoa(i), ansMap[i])
	}
	return ans
}
