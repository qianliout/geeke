package main

import (
	"fmt"
)

func main() {
	next := GetLastNext("abab")
	fmt.Println("res is ", next)
}

/*
给定一个字符串 s，你可以通过在字符串前面添加字符将其转换为回文串。找到并返回可以用这种方式转换的最短回文串。
示例 1:
输入: "aacecaaa"
输出: "aaacecaaa"
示例 2:
输入: "abcd"
输出: "dcbabcd"
*/

func shortestPalindrome(s string) string {
	return ""
}

/*
    //返回 next 数组的最后一个值
public int getLastNext(String s) {
    int n = s.length();
    char[] c = s.toCharArray();
    int[] next = new int[n + 1];
    next[0] = -1;
    next[1] = 0;
    int k = 0;
    int i = 2;
    while (i <= n) {
        if (k == -1 || c[i - 1] == c[k]) {
            next[i] = k + 1;
            k++;
            i++;
        } else {
            k = next[k];
        }
    }
    return next[n];
}
*/

func GetLastNext(s string) []int {
	n := len(s)
	ss := []byte(s)
	next := make([]int, n+1)
	next[0] = -1
	next[1] = 0
	k, i := 0, 2
	for i <= n {
		if k == -1 || ss[i-1] == ss[k] {
			next[i] = k + 1
			k++
			i++
		} else {
			k = next[k]
		}
	}
	return next
}
