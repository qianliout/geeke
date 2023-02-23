package main

import (
	"fmt"
	"math"
)

func main() {
	res := integerReplacement(7)
	fmt.Println("res is ", res)
	fmt.Println(7 & 2)
}

/*
给定一个正整数 n，你可以做如下操作：
1. 如果 n 是偶数，则用 n / 2替换 n。
2. 如果 n 是奇数，则可以用 n + 1或n - 1替换 n。
n 变为 1 所需的最小替换次数是多少？
示例 1:
输入:
8
输出:
3
解释:
8 -> 4 -> 2 -> 1
示例 2:
输入:
7
输出:
4
解释:
7 -> 8 -> 4 -> 2 -> 1
或
7 -> 6 -> 3 -> 2 -> 1
*/
func integerReplacement(n int) int {
	if n <= 1 {
		return 0
	}
	if n <= 3 {
		return n - 1
	}

	res := 0
	for n > 1 {
		if n&1 == 0 {
			n >>= 1
			res++
		} else {
			if n == 3 {
				res++
				n -= 1
			} else if n&2 == 0 { // 这里用位操作纯粹是炫技,目的就是 让每一步1的数目最少好处大，于是 0bxxx01 采用 -1； 0bxxx11 采用 +1；
				res++
				n -= 1
			} else {
				res++
				n += 1
			}
		}
	}
	return res
}

// 这道题主要处理3,第二点,就是尽量往4上靠
func integerReplacement2(n int) int {
	if n <= 1 {
		return 0
	}
	if n == math.MaxInt32 {
		return 32
	}
	if n <= 3 {
		return n - 1
	}
	if n%2 == 0 {
		return integerReplacement2(n/2) + 1
	} else if n%4 == 1 {
		return integerReplacement2(n-1) + 1
	} else if n%4 == 3 {
		return integerReplacement2(n+1) + 1
	}
	return 0
}

func integerReplacement3(n int) int {
	if n <= 1 {
		return 0
	}
	res := 0
	for n > 1 {
		if n <= 3 {
			res = res + n - 1
			break
		}
		if n%4 == 3 {
			res++
			n = n + 1
		}
		if n%4 == 1 {
			res++
			n = n - 1
		}
		if n%2 == 0 {
			res++
			n = n / 2
		}
	}
	return res
}
