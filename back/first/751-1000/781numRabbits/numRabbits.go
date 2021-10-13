package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numRabbits([]int{1,1,1,1,2}))
}

/*
森林中，每个兔子都有颜色。其中一些兔子（可能是全部）告诉你还有多少其他的兔子和自己有相同的颜色。我们将这些回答放在 answers 数组里。
返回森林中兔子的最少数量。
示例:
输入: answers = [1, 1, 2]
输出: 5
解释:
两只回答了 "1" 的兔子可能有相同的颜色，设为红色。
之后回答了 "2" 的兔子不会是红色，否则他们的回答会相互矛盾。
设回答了 "2" 的兔子为蓝色。
此外，森林中还应有另外 2 只蓝色兔子的回答没有包含在数组中。
因此森林中兔子的最少数量是 5: 3 只回答的和 2 只没有回答的。
输入: answers = [10, 10, 10]
输出: 11
输入: answers = []
输出: 0
说明:
answers 的长度最大为1000。
answers[i] 是在 [0, 999] 范围内的整数。
*/
func numRabbits(answers []int) int {
	sort.Ints(answers)
	
	ans, i := 0, 0
	for i < len(answers) {
		ans += answers[i] + 1
		start, k := answers[i], answers[i]
		
		i++
		// 还可以向生走多少个
		for k > 0 && i < len(answers) && answers[i] == start {
			k--
			i++
		}
	}
	return ans
}
