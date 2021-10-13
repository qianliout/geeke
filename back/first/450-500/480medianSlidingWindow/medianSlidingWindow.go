package main

import (
	"container/heap"
	"fmt"

	"outback/leetcode/back/common/commonHeap"
)

func main() {
	// nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	// nums := []int{1, 2, 3, 4, 2, 3, 1, 4, 2}
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7} // [1.0,-1.0,-1.0,3.0,5.0,6.0]
	window := medianSlidingWindow(nums, 3)
	fmt.Println("res is ", window)
}

/*
中位数是有序序列最中间的那个数。如果序列的大小是偶数，则没有最中间的数；此时中位数是最中间的两个数的平均数。
例如：
[2,3,4]，中位数是 3
[2,3]，中位数是 (2 + 3) / 2 = 2.5
给你一个数组 nums，有一个大小为 k 的窗口从最左端滑动到最右端。窗口中有 k 个数，每次窗口向右移动 1 位。
你的任务是找出每次窗口移动后得到的新窗口中元素的中位数，并输出由它们组成的数组。
示例：
给出 nums = [1,3,-1,-3,5,3,6,7]，以及 k = 3。
窗口位置                      中位数
---------------               -----
[1  3  -1] -3  5  3  6  7       1
 1 [3  -1  -3] 5  3  6  7      -1
 1  3 [-1  -3  5] 3  6  7      -1
 1  3  -1 [-3  5  3] 6  7       3
 1  3  -1  -3 [5  3  6] 7       5
 1  3  -1  -3  5 [3  6  7]      6
 因此，返回该滑动窗口的中位数数组 [1,-1,-1,3,5,6]。
提示：
你可以假设 k 始终有效，即：k 始终小于输入的非空数组的元素个数。
与真实值误差在 10 ^ -5 以内的答案将被视作正确答案。
*/

// 用两个堆,如果k是奇数,中位数在小顶堆上,如果k中偶数,值是两堆的平均数
func medianSlidingWindow(nums []int, k int) []float64 {
	res := make([]float64, 0)
	if len(nums) < k {
		return res
	}
	// 先构建初始堆
	rightHeap := make(commonHeap.MinHeap, 0)
	leftHeap := make(commonHeap.MaxHeap, 0)
	length := len(nums)
	deleteMap := make(map[int]int)
	// 用这两个变量来控制堆的平衡
	leftNodeNum, rightNodeNum := 0, 0

	// 先初始两个堆,先放在小堆,也就是后面的堆里
	for i := 0; i < k; i++ {
		heap.Push(&rightHeap, nums[i])
		rightNodeNum++
	}
	for j := 0; j < (k+1)/2; j++ {
		min := heap.Pop(&rightHeap).(int)
		rightNodeNum--
		heap.Push(&leftHeap, min)
		leftNodeNum++
	}
	// fmt.Println("init",leftHeap, rightHeap)
	for i := k; i < length; i++ {
		// 先把中位数放在结果集中,因为这里是先加结果,所以在最后一次循环是,还有一组结果没有加,所以要加上
		if k%2 == 0 {
			res = append(res, float64(rightHeap[0]+leftHeap[0])/2)
		} else {
			res = append(res, float64(leftHeap[0]))
		}

		outNum, inNum := nums[i-k], nums[i]

		// 移出元素判断是在那一边
		if outNum <= leftHeap[0] {
			leftNodeNum--
		} else {
			rightNodeNum--
		}
		deleteMap[outNum]++
		// 加入元素
		if len(leftHeap) > 0 && inNum <= leftHeap[0] {
			heap.Push(&leftHeap, inNum)
			leftNodeNum++
		} else {
			heap.Push(&rightHeap, inNum)
			rightNodeNum++
		}
		// 加入元素后重新平衡
		// fmt.Println("111111", nums[i], leftHeap, rightHeap, leftNodeNum, rightNodeNum)
		// 因为我的设定是,如果是奇数,左边比右边多一个,所以右边是不能多于左边的
		if rightNodeNum > leftNodeNum {
			heap.Push(&leftHeap, heap.Pop(&rightHeap).(int))
			rightNodeNum--
			leftNodeNum++
		}
		// 左边比右边多的数大于一了
		if leftNodeNum-rightNodeNum > 1 {
			heap.Push(&rightHeap, heap.Pop(&leftHeap).(int))
			rightNodeNum++
			leftNodeNum--
		}
		// fmt.Println("2222222", leftHeap, rightHeap)
		// 延迟删除
		for len(leftHeap) > 0 {
			if deleteMap[leftHeap[0]] > 0 {
				p := heap.Pop(&leftHeap).(int)
				deleteMap[p]--
			} else {
				break
			}
		}
		for len(rightHeap) > 0 {
			if deleteMap[rightHeap[0]] > 0 {
				p := heap.Pop(&rightHeap).(int)
				deleteMap[p]--
			} else {
				break
			}
		}
	}

	// 这里要注意,最后一条记录容易忘记
	if k%2 == 0 {
		res = append(res, float64(rightHeap[0]+leftHeap[0])/2)
	} else {
		res = append(res, float64(leftHeap[0]))
	}
	return res
}
