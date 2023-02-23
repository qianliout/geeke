package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums := []int{1, 2, 3}
	s := Constructor(nums)
	fmt.Println(s.Shuffle())
	fmt.Println(s.Shuffle())
	fmt.Println(s.Shuffle())
	fmt.Println(s.Reset())
}

/*
打乱一个没有重复元素的数组。
示例:
// 以数字集合 1, 2 和 3 初始化数组。
int[] nums = {1,2,3};
Solution solution = new Solution(nums);
// 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。
solution.shuffle();
// 重设数组到它的初始状态[1,2,3]。
solution.reset();
// 随机返回数组[1,2,3]打乱后的结果。
solution.shuffle();
*/

type Solution struct {
	nums1 []int // 原数据
	nums2 []int // 上一次洗牌后的数组
	rand  *rand.Rand
}

func Constructor(nums []int) Solution {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	return Solution{
		nums1: append([]int{}, nums...),
		nums2: append([]int{}, nums...),
		rand:  r,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	this.rand = rand.New(rand.NewSource(time.Now().Unix()))
	this.nums2 = append([]int{}, this.nums1...)
	return this.nums1
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	n := len(this.nums2)
	for i := 0; i < n; i++ {
		index := rand.Intn(n)
		this.nums2[i], this.nums2[index] = this.nums2[index], this.nums2[i]
	}

	return this.nums2
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
