package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 4, 6, 2}
	//ints := []int{4, 5, 6}
	//h := append(nums[:2], nums[3:]...)
	//fmt.Println(len(h))
	res := permute6(nums)
	fmt.Println(res)
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	var backtrck func(nums, tep []int)

	backtrck = func(ints, tep []int) {
		if len(ints) == 0 {
			res = append(res, tep)
			return
		}
		for i := 0; i < len(ints); i++ {

			tep = append(tep, ints[i])
			back := append(ints[:i], ints[i+1:]...)
			backtrck(back, tep)
		}
	}
	tem := make([]int, 0)
	backtrck(nums, tem)
	return res
}

func permute1(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	ret := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		buf := make([]int, len(nums)-1)
		copy(buf, nums[0:i])
		copy(buf[i:], nums[i+1:])
		r := permute(buf)
		for _, v := range r {
			ret = append(ret, append(v, nums[i]))
		}
	}
	return ret
}

func permute3(nums []int) [][]int {
	res := make([][]int, 0)
	length := len(nums)

	var backtrack func(first int)

	backtrack = func(first int) {
		if first >= length {
			res = append(res, nums)
			return
		}
		for i := first; i < length; i++ {
			//	sawp
			a := nums[first]
			b := nums[i]
			nums[first] = b
			nums[i] = a
			fmt.Println("swap nums first:", nums)
			backtrack(first + 1)

			c := nums[first]
			d := nums[i]
			nums[first] = d
			nums[i] = c
			//fmt.Println("swap nums seconde:", nums)
		}
	}

	backtrack(0)
	return res
}

//这里声明一个全局变量用来存储所有的排列
var result [][]int

func permute4(nums []int) [][]int {

	//每次调用重置result全局变量，防止结果缓存
	result = make([][]int, 0, 2*len(nums))

	//当nums只有一个元素的情况下，直接返回即可
	if len(nums) == 1 {
		result = append(result, nums)
		return result
	}

	//声明一个arr变量，用来存储路径
	arr := make([]int, 0, len(nums))

	arrange(nums, arr)

	return result
}

func arrange(nums []int, arr []int) {

	//当nums长度为0，选择列表为空，路径选择完毕，返回即可
	if len(nums) == 0 {
		return
	}

	//循环当前的选择列表
	for k, v := range nums {
		//选取一个元素，例如选1，则使用newArr来存储更新后的选择列表
		arr = append(arr, v)
		newArr := make([]int, len(nums)-1)
		copy(newArr[:k], nums[:k])
		if k < len(nums)-1 {
			copy(newArr[k:], nums[k+1:])
		}
		//递归调用，将更新后的选择列表和存储路径的arr传入
		arrange(newArr, arr)
		//当arr的长度和容量相等，说明路径已经选择完毕
		//将此条路径记录到结果中
		if len(arr) == cap(arr) {
			path := make([]int, len(arr))
			copy(path, arr)
			result = append(result, path)
		}
		/**
		返回上一个路口重新做选择
		例如，当arr为1开头的时候，下一个路口有2，3两种选择
		当选择完毕，将路径记录之后，需要返回重新选择
		例如从1->2 返回到 1->3的结果
		**/
		arr = arr[:len(arr)-1]
	}
}
