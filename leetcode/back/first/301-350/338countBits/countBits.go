package main

import (
	"fmt"
)

func main() {
	res := countBits(2)
	fmt.Println("res is ", res)
}

/*

给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。
示例 1:
输入: 2
输出: [0,1,1]
示例 2:
输入: 5
输出: [0,1,1,2,1,2]
进阶:
    给出时间复杂度为O(n*sizeof(integer))的解答非常容易。但你可以在线性时间O(n)内用一趟扫描做到吗？
    要求算法的空间复杂度为O(n)。
    你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的 __builtin_popcount）来执行此操作。
链接：https://leetcode-cn.com/problems/counting-bits
*/
// 首先想到的是循环计算
// 另一种方法是dp
func countBits(num int) []int {
	// return find1(num)
	return find3(num)
}

// dp 在已计算好的数字后面加一个1
// dp[x] = dp[x/2]+ x%2 如果x是偶数，二倍只是在最后面加1，如果是奇数，x/2 则还要在最后面加1
func find2(num int) []int {
	res := make([]int, 0)
	dp := make(map[int]int)
	dp[1], dp[0] = 1, 0
	for i := 1; i <= num; i++ {
		dp[i] = dp[i>>1] + i&1
		res = append(res, dp[i])
	}
	return res
}

// dp 通过在高位加1的方法，虽然该方法也可以，但是不直观
// 方法四：动态规划 + 最后设置位【通过】
// 最后设置位是从右到左第一个为1的位。使用 x &= x - 1 将该位设置为0，就可以得到以下状态转移函数：
// P(x)=P(x&(x−1))+1
func find3(num int) []int {
	res := make([]int, 0)
	dp := make(map[int]int)
	dp[1], dp[0] = 1, 0
	res = append(res, 0)
	for i := 1; i <= num; i++ {
		dp[i] = dp[i&(i-1)] + 1
		res = append(res, dp[i])
	}
	return res
}

func find1(num int) []int {
	res := make([]int, 0)
	for i := 0; i <= num; i++ {
		res = append(res, huming(i))
	}
	return res
}

func huming(num int) int {
	count := 0
	for num > 0 {
		if num&1 == 1 {
			count++
		}
		num >>= 1
	}
	return count
}
