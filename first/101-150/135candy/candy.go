package main

import (
	"fmt"
	"math"
)

func main() {
	rating := []int{1, 2, 2}
	res := candy(rating)
	fmt.Println("res is ", res)
}

/*
老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。
你需要按照以下要求，帮助老师给这些孩子分发糖果：
    每个孩子至少分配到 1 个糖果。
    相邻的孩子中，评分高的孩子必须获得更多的糖果。
那么这样下来，老师至少需要准备多少颗糖果呢？
示例 1:
输入: [1,0,2]
输出: 5
解释: 你可以分别给这三个孩子分发 2、1、2 颗糖果。
示例 2:
输入: [1,2,2]
输出: 4
解释: 你可以分别给这三个孩子分发 1、2、1 颗糖果。
     第三个孩子只得到 1 颗糖果，这已满足上述两个条件。
链接：https://leetcode-cn.com/problems/candy
*/
func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	n := len(ratings)
	left2right := make([]int, n)
	right2left := make([]int, n)
	for i := 0; i < n; i++ {
		left2right[i] = 1
		if i >= 1 && ratings[i] > ratings[i-1] {
			left2right[i] = left2right[i-1] + 1
		}
	}
	// 这里可以有用再开一个数据，但是为了好理解，这里使用这种写法
	for i := n - 1; i >= 0; i-- {
		right2left[i] = 1
		if i < n-1 && ratings[i] > ratings[i+1] {
			right2left[i] = right2left[i+1] + 1
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		res += int(math.Max(float64(left2right[i]), float64(right2left[i])))
	}
	return res
}
