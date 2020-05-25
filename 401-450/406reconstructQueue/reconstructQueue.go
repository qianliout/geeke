package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	res := reconstructQueue(nums)
	fmt.Println("res is ", res)
}

/*
假设有打乱顺序的一群人站成一个队列。 每个人由一个整数对(h, k)表示，其中h是这个人的身高，
k是排在这个人前面且身高大于或等于h的人数。 编写一个算法来重建这个队列。
注意：
总人数少于1100人。
示例
输入:
因为个子矮的人相对于个子高的人是看不见的。。。是看不见的。。。不见的。。。
[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
输出:
[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
*/

// 其实就是一个排序:因为个子矮的人相对于个子高的人是看不见的,这个很重要
func reconstructQueue(people [][]int) [][]int {
	sort.Sort(Item(people))
	// 按要求放值 [[7 0] [7 1] [6 1] [5 0] [5 2] [4 4]]
	res := make([][]int, 0)
	for _, value := range people {
		res = insertFunc(res, value[1], value)
	}
	return res
}

// 随机插入的辅助函数，这个方法的复杂度其实是n方的
func insertFunc(nums [][]int, pos int, target []int) [][]int {
	tmp := append([][]int{}, nums[pos:]...)
	if pos >= len(nums) {
		res := append(nums, target)
		return res
	} else {
		nums = nums[:pos+1]
		nums[pos] = target
		res := append(nums, tmp...)
		return res
	}
}

type Item [][]int

func (it Item) Len() int {
	return len(it)
}

func (it Item) Less(i, j int) bool {
	if it[i][0] > it[j][0] {
		return true
	} else if it[i][0] < it[j][0] {
		return false
	} else {
		if it[i][1] < it[j][1] {
			return true
		} else {
			return false
		}
	}
}

func (it Item) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}
