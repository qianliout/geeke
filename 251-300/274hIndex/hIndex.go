package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{3, 0, 6, 1, 5}

	ints := Sort(nums)
	fmt.Println("sort is ", ints)
	//nums := []int{100}
	res := hIndex3(nums)
	fmt.Println("res is ", res)
}

/*
给定一位研究者论文被引用次数的数组（被引用次数是非负整数）。编写一个方法，计算出研究者的 h 指数。
h 指数的定义: “h 代表“高引用次数”（high citations），一名科研人员的 h 指数是指他（她）的 （N 篇论文中）
至多有 h 篇论文分别被引用了至少 h 次。（其余的 N - h 篇论文每篇被引用次数不多于 h 次。）”
示例:
输入: citations = [3,0,6,1,5]
输出: 3
解释: 给定数组表示研究者总共有 5 篇论文，每篇论文相应的被引用了 3, 0, 6, 1, 5 次。
     由于研究者有 3 篇论文每篇至少被引用了 3 次，其余两篇论文每篇被引用不多于 3 次，所以她的 h 指数是 3。
链接：https://leetcode-cn.com/problems/h-index
*/

func hIndex(citations []int) int {
	if len(citations) == 0 || (len(citations) == 1 && citations[0] == 0) {
		return 0
	} else if len(citations) == 1 {
		return 1
	}

	sort.Ints(citations)
	length := len(citations)
	for i := 0; i < len(citations); i++ {
		if citations[length-i-1] > i {
			continue
		}
		return i
	}
	//如果把整个数据查找完了，都还是满足，就返回数据长度
	return length
}

// 本题可以只关心计数，而不用真正的排序
func hIndex3(citations []int) int {
	length := len(citations)

	tem := make([]int, length+1)
	for _, n := range citations {
		tem[min(length, n)]++
	}
	fmt.Println("tem ", tem)
	// 这里有点不懂
	k := length
	s := tem[length]
	for k > s {
		k--
		s += tem[k] // 表示一共出现了几个论文了
	}
	return k
}

func min(nums ...int) int {
	m := math.MaxInt64
	for _, n := range nums {
		if n < m {
			m = n
		}
	}
	return m
}

func hIndex2(citations []int) int {
	sort.Ints(citations)
	i := 0

	for i < len(citations) && citations[len(citations)-i-1] > i {
		i++
	}
	return i
}

// 上述排序是一般的比较排序，时间复杂底是O（nlog(n)）,可以使用计数排序，时间复杂度更低
func Sort(nums []int) []int {
	min, max := math.MaxInt64, math.MinInt64
	// 每一步确定最大最小值
	for _, i := range nums {
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}
	tem := make([]int, max-min+1)
	for _, i := range nums {
		tem[i-min] += 1
	}
	res := make([]int, len(nums))
	j := 0
	for i, _ := range tem {
		for tem[i] > 0 {
			res[j] = i + min
			tem[i]--
			j++
		}
	}
	return res
}
