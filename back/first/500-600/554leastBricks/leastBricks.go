package main

import (
	"fmt"

	"qianliout/leetcode/common/utils"
)

func main() {
	wail := [][]int{{1, 2, 2, 1}, {3, 1, 2}, {1, 3, 2}, {2, 4}, {3, 1, 2}, {1, 3, 1, 1}}
	res := leastBricks(wail)
	fmt.Println("res is ", res)
}

/*
你的面前有一堵矩形的、由多行砖块组成的砖墙。 这些砖块高度相同但是宽度不同。你现在要画一条自顶向下的、穿过最少砖块的垂线。
砖墙由行的列表表示。 每一行都是一个代表从左至右每块砖的宽度的整数列表。
如果你画的线只是从砖块的边缘经过，就不算穿过这块砖。你需要找出怎样画才能使这条线穿过的砖块数量最少，并且返回穿过的砖块数量。
你不能沿着墙的两个垂直边缘之一画线，这样显然是没有穿过一块砖的。
示例：
输入: [[1,2,2,1],
      [3,1,2],
      [1,3,2],
      [2,4],
      [3,1,2],
      [1,3,1,1]]
输出: 2
解释:
提示：
    每一行砖块的宽度之和应该相等，并且不能超过 INT_MAX。
    每一行砖块的数量在 [1,10,000] 范围内， 墙的高度在 [1,10,000] 范围内， 总的砖块数量不超过 20,000。
*/

// 前缀和的用法
func leastBricks(wall [][]int) int {

	sumMap := make(map[int]int)
	ans := 0

	for _, list := range wall {
		sum := 0
		// 开始和最后要排除
		for i := 0; i < len(list)-1; i++ {
			sum += list[i]
			if k, ok := sumMap[sum]; ok {
				sumMap[sum] = k + 1
			} else {
				sumMap[sum] = 1
			}
			ans = utils.Max(ans, sumMap[sum])
		}
	}
	return len(wall) - ans
}
