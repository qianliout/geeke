package main

import (
	"fmt"
	"math"
)

func main() {
	res := grayCode2(3)
	fmt.Println("res is ", res)
}

/*
格雷编码是一个二进制数字系统，在该系统中，两个连续的数值仅有一个位数的差异。
给定一个代表编码总位数的非负整数 n，打印其格雷编码序列。格雷编码序列必须以 0 开头。
示例 1:
输入: 2
输出: [0,1,3,2]
解释:
00 - 0
01 - 1
11 - 3
10 - 2
对于给定的 n，其格雷编码序列并不唯一。
例如，[0,2,3,1] 也是一个有效的格雷编码序列。
00 - 0
10 - 2
11 - 3
01 - 1
示例 2:
输入: 0
输出: [0]
解释: 我们定义格雷编码序列必须以 0 开头。
     给定编码总位数为 n 的格雷编码序列，其长度为 2n。当 n = 0 时，长度为 20 = 1。
     因此，当 n = 0 时，其格雷编码序列为 [0]。
*/

// 可以使用镜面的方式，这里可以使用两个stark，但是也可以使用一个stark
func grayCode(n int) []int {
	firstStack := make([]int, 0)
	//secondStack := make([]int, 0)
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}
	firstStack = append(firstStack, 0, 1)
	for i := 2; i <= n; i++ {
		//先正的放
		//for _, j := range firstStack {
		//	secondStack = append(secondStack, j)
		//}
		// 再镜面放
		for j := len(firstStack) - 1; j >= 0; j-- {
			newNum := int(math.Pow(float64(2), float64(i-1))) + firstStack[j]
			firstStack = append(firstStack, newNum)
		}

		//firstStack = secondStack
		//secondStack = make([]int, 0)
	}
	return firstStack
}

// 直接使用公式法
//二进制转成格雷码有一个公式。
//所以我们遍历 0 到 2n−12^n-12n−1，然后利用公式转换即可。即最高位保留，其它位是当前位和它的高一位进行异或操作。

func grayCode2(n int) []int {
	res := make([]int, 0)

	for i := 0; i < 1<<n; i++ {
		res = append(res, i^i>>1)
	}
	return res
}
