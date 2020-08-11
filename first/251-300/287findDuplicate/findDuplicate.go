package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 4, 2, 2}
	res := findDuplicate(nums)
	fmt.Println("res is ", res)
}

/*
给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），
可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。
示例 1:
输入: [1,3,4,2,2]
输出: 2
示例 2:
输入: [3,1,3,4,2]
输出: 3
说明：
    不能更改原数组（假设数组是只读的）。
    只能使用额外的 O(1) 的空间。
    时间复杂度小于 O(n2) 。
    数组中只有一个重复的数字，但它可能不止重复出现一次。
*/
// 可以使用map，排序的方法，但是不满足题目中的要求
//所以只能使用快慢指针的方式，也叫乌龟和兔子的方法,和判断链表中是否有环是一个道理
func findDuplicate(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow, fast, find := 0, 0, 0
	for {
		slow, fast = nums[slow], nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	for find != slow {
		find, slow = nums[find], nums[slow]
	}
	return find
}
