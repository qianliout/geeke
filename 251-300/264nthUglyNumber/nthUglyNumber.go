package main

import (
	"fmt"
	"math"
)

func main() {
	res := nthUglyNumber(2)
	fmt.Println("res is ", res)
}

/*编写一个程序，找出第 n 个丑数。
丑数就是只包含质因数 2, 3, 5 的正整数。
示例:
输入: n = 10
输出: 12
解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
说明:
    1 是丑数。
    n 不超过1690。
*/
// 方法一动态规划
func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	i2, i3, i5 := 0, 0, 0
	nums := make([]int, 0)
	exitmap := make(map[int]bool)
	nums = append(nums, 1)
	i := 0
	for {
		tem := minInt(nums[i2]*2, nums[i3]*3, nums[i5]*5)
		if !exitmap[tem] {
			nums = append(nums, tem)
			exitmap[tem] = true
		}
		if tem == nums[i2]*2 {
			i2++
		} else if tem == nums[i3]*3 {
			i3++
		} else if tem == nums[i5]*5 {
			i5++
		}
		i++
		if len(nums) > n+1 {
			break
		}
	}
	return nums[n-1]
}

func minInt(nums ...int) int {
	res := math.MaxInt64
	for _, i2 := range nums {
		if i2 < res {
			res = i2
		}
	}
	return res
}
