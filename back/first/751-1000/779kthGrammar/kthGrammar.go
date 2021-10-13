package main

import (
	"math"
)

func main() {
	
}

/*
在第一行我们写上一个 0。接下来的每一行，将前一行中的0替换为01，1替换为10。
给定行数 N 和序数 K，返回第 N 行中第 K个字符。（K从1开始）
例子:
输入: N = 1, K = 1
输出: 0
输入: N = 2, K = 1
输出: 0
输入: N = 2, K = 2
输出: 1
输入: N = 4, K = 5
输出: 1
解释:
第一行: 0
第二行: 01
第三行: 0110
第四行: 01101001
*/
/*
直接找规律。
第一行 0
第二行 01
第三行 0110
第四行 01101001
可以发现，第n行的数量比第n-1行多了一倍，并且前半部分是和第n-1行一样的，后半部分是前半部分“按位取反”得到的。
第n行的字符数量是2^(n-1)个，因此第n-1行的数量就是2^(n-2)个。公式为：
func(n,k) = func(n-1,k), if k <= 2^(n-2)
func(n,k)= ^func(n-1,k-2^(n-2)), if k >2^(n-2)
*/
func kthGrammar(N int, K int) int {
	if N == 1 || N == 0 {
		return 0
	}
	tmp := int(math.Pow(2, float64(N-2)))
	if K < tmp {
		return kthGrammar(N-1, K)
	} else {
		r := kthGrammar(N-1, K-tmp)
		if r == 1 {
			return 0
		} else {
			return 1
		}
	}
}
