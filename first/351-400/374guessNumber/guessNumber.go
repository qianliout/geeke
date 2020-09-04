package main

import (
	"fmt"
)

func main() {
	res := guessNumber(1)
	fmt.Println(res)
}

/*
猜数字游戏的规则如下：
    每轮游戏，系统都会从 1 到 n 随机选择一个数字。 请你猜选出的是哪个数字。
    如果你猜错了，系统会告诉你，你猜测的数字比系统选出的数字是大了还是小了。
你可以通过调用一个预先定义好的接口 guess(int num) 来获取猜测结果，返回值一共有 3 种可能的情况（-1，1 或 0）：
-1 : 你猜测的数字比系统选出的数字大
 1 : 你猜测的数字比系统选出的数字小
 0 : 恭喜！你猜对了！
示例 :
输入: n = 10, pick = 6
输出: 6
*/

/*
int mid = (left + right) / 2; 是初级写法，是有 bug 的；
2、int mid = left + (right - left) / 2; 是正确的写法，说明你考虑到了整型溢出的风险,但是当right是大的正数，
而left是一个很小的负数时也会有溢出的可能
3、int mid = (low + high) >>> 1; 无符号位移 首先肯定是正确的写法，其实也是一个装 ❌ 的写法，。go语言不支持
*/
func guessNumber(n int) int {
	left, right := 1, n
	
	for left <= right {
		mid := left + (right-left)/2
		r := guess(mid, 1)
		if r == 0 {
			return mid
		} else if r == 0 {
			left = mid + 1
		} else if r == -1 {
			right = mid - 1
		}
	}
	return 0
}

func guess(n, peek int) int {
	if n == peek {
		return 0
	} else if n > peek {
		return 1
	} else {
		return -1
	}
}
