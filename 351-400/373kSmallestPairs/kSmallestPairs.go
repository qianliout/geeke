package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	pairs := kSmallestPairs(nums1, nums2, 3)
	fmt.Println("res is ", pairs)
}

/*
给定两个以升序排列的整形数组 nums1 和 nums2, 以及一个整数 k。
定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2。
找到和最小的 k 对数字 (u1,v1), (u2,v2) ... (uk,vk)。
示例 1:
输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
输出: [1,2],[1,4],[1,6]
解释: 返回序列中的前 3 对数：
     [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
示例 2:
输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
输出: [1,1],[1,1]
解释: 返回序列中的前 2 对数：
     [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
示例 3:
输入: nums1 = [1,2], nums2 = [3], k = 3
输出: [1,3],[2,3]
解释: 也可能序列中所有的数对都被返回:[1,3],[2,3]
*/

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	ans := make(MaxHeap, 0)

	if k == 0 {
		return [][]int{}
	}
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			if len(ans) < k {
				heap.Push(&ans, Item{v1, v2})
			} else {

				top := ans[0]
				//fmt.Println("top is ", top, ans)
				if v1+v2 >= top[0]+top[1] {
					continue
				} else {
					//fmt.Println("push ", v1, v2, ans)
					heap.Pop(&ans)
					heap.Push(&ans, Item{v1, v2})
				}
			}
		}
	}
	fmt.Println(ans)
	answer := make([]Item, 0)
	for len(ans) > 0 {
		answer = append(answer, heap.Pop(&ans).(Item))
		heap.Init(&ans)
	}
	answer2 := make([][]int, 0)
	for i := len(answer) - 1; i >= 0; i-- {
		answer2 = append(answer2, []int{answer[i][0],answer[i][1]})
	}
	return answer2
}
                         
type Item []int

type MaxHeap []Item

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i][0]+h[i][1] > h[j][0]+h[j][1] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
