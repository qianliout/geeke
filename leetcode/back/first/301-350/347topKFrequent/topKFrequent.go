package main

import (
	"container/heap"
	"fmt"

	"outback/leetcode/back/common/commonHeap"
)

func main() {
	num := []int{3, 0, 1, 0, 1, 1, 2}
	res := topKFrequent(num, 1)
	fmt.Println("res is ", res)
}

/*
347. 前 K 个高频元素
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
示例 1:
输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:
输入: nums = [1], k = 1
输出: [1]
说明：
    你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
    你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
*/
// 使用优先队列
func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0)
	if len(nums) == 0 || k == 0 {
		return res
	}
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num] += 1
	}
	// fmt.Println(frequency)
	// 构建优先队列
	pq := make(commonHeap.PriorityQueue, 0)
	for i, v := range frequency {
		heap.Push(&pq, &commonHeap.IntItem{Value: i, Priority: v})
	}
	// for _, v := range pq {
	//	fmt.Println(v.Key, v.Priority)
	// }

	for i := 0; i < k; i++ {
		// value:=heap.Pop(&pq).(*IntItem).Value
		value := heap.Remove(&pq, 0).(*commonHeap.IntItem).Value
		res = append(res, value)
	}
	return res
}
