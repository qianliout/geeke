package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s := Constructor([]int{1, 2, 3, 3, 3})
	fmt.Println(s.IndexMap)
	fmt.Println(s.Frequency)
	res := s.Pick(3)
	fmt.Println("res is ", res)
}

/*
给定一个可能含有重复元素的整数数组，要求随机输出给定的数字的索引。 您可以假设给定的数字一定存在于数组中。
注意：
数组大小可能非常大。 使用太多额外空间的解决方案将不会通过测试。
示例:
int[] nums = new int[] {1,2,3,3,3};
Solution solution = new Solution(nums);
// pick(3) 应该返回索引 2,3 或者 4。每个索引的返回概率应该相等。
solution.pick(3);
// pick(1) 应该返回 0。因为只有nums[0]等于1。
solution.pick(1);
*/

type Solution struct {
	Frequency map[int]int
	IndexMap  map[int][]int
	Nums      []int
	Rand      *rand.Rand
}

func Constructor(nums []int) Solution {
	f := make(map[int]int)
	indexMap := make(map[int][]int)
	for i, num := range nums {
		f[num]++
		if _, ok := indexMap[num]; !ok {
			indexMap[num] = make([]int, 0)
		}
		indexMap[num] = append(indexMap[num], i)
	}
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	return Solution{Nums: nums, Frequency: f, Rand: r, IndexMap: indexMap}
}

func (this *Solution) Pick(target int) int {
	v, ok := this.Frequency[target]
	if !ok {
		return -1
	}
	iMap := this.IndexMap[target]
	ans := iMap[0]
	for i := 1; i <= v; i++ {
		if i == this.Rand.Intn(i)+1 {
			ans = iMap[i-1]
		}
	}
	return ans
}
