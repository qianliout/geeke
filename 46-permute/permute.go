package main

//func permute5(nums []int) [][]int {
//	res = make([][]int, 0)
//	var backtrack func(nums, tep []int)
//
//	backtrack = func(ints, tep []int) {
//		if len(ints) == 0 {
//			res = append(res, tep)
//			return
//		}
//		for i := 0; i < len(ints); i++ {
//			tep = append(tep, ints[i])
//			ints := append(ints[:i], ints[i+1:]...)
//			backtrack(ints, tep)
//		}
//	}
//	tem := make([]int, 0)
//	backtrack(nums, tem)
//	return res
//}
var res [][]int

func permute6(nums []int) [][]int {

	res = make([][]int, 0)
	var tmp []int
	//visited := make([]int, len(nums))
	//backtrack(nums, tmp, visited)
	used := make([]bool, len(nums))
	dfs(nums, len(nums), 0, tmp, used)

	return res
}

func backtrack(nums, tmp, visited []int) {
	if len(tmp) == len(nums) {
		//fmt.Println("tmp ", tmp)
		res = append(res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] == 1 {
			continue
		}
		visited[i] = 1
		tmp = append(tmp, nums[i])
		backtrack(nums, tmp, visited)
		visited[i] = 0
		tmp = tmp[:len(tmp)-1]
	}
}

func dfs(nums []int, length, depth int, path []int, used []bool) {
	if depth == length {
		res = append(res, path)
		return
	}
	for i := 0; i < length; i++ {
		if !used[i] {

			used[i] = true
			path = append(path, nums[i])
			dfs(nums, length, depth+1, path, used)
			used[i] = false
			//path = path[:len(path)-1]
			path = append(path[:depth], path[depth+1:]...)
		}
	}
}
