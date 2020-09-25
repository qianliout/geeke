package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	houses := []int{1, 2, 3, 20}
	heaters := []int{1, 4}
	res := findRadius(houses, heaters)
	fmt.Println("res is ", res)
}

/*
冬季已经来临。 你的任务是设计一个有固定加热半径的供暖器向所有房屋供暖。
现在，给出位于一条水平线上的房屋和供暖器的位置，找到可以覆盖所有房屋的最小加热半径。
所以，你的输入将会是房屋和供暖器的位置。你将输出供暖器的最小加热半径。
说明:
    给出的房屋和供暖器的数目是非负数且不会超过 25000。
    给出的房屋和供暖器的位置均是非负数且不会超过10^9。
    只要房屋位于供暖器的半径内(包括在边缘上)，它就可以得到供暖。
    所有供暖器都遵循你的半径标准，加热的半径也一样。
示例 1:
输入: [1,2,3],[2]
输出: 1
解释: 仅在位置2上有一个供暖器。如果我们将加热半径设为1，那么所有房屋就都能得到供暖。
示例 2:
输入: [1,2,3,4],[1,4]
输出: 1
解释: 在位置1, 4上有两个供暖器。我们需要将加热半径设为1，这样所有房屋就都能得到供暖。
*/

// 第一步以每一个房屋为研究中心，然后遍历加热站（可以利用二分法节省效率）找到房屋的前一个加热站和后一个加热站。
// 接着通过比较房屋和前后加热站的距离比较房屋和谁更近。
//
// 对于所有房屋h(1),h(2),h(3)···h(n),每个房屋与最近的加热站的距离为dis(n),n为房间编号。即每个房屋距离加热站
// 的最短距离disList为：
// disList=[dis(1),dis(2),dis(3)···dis(n)]
// 因此：
// 当取disList中的最大值dis-max时。一定满足dis-max大于等于disList中的每一项。即dis-max一定能够辐射到每一个房屋。
// 为此dis-max即为问题的解。
// 二分查找
func findRadius(houses []int, heaters []int) int {
	if len(heaters) == 0 || len(heaters) == 0 {
		return 0
	}
	sort.Ints(heaters)
	sort.Ints(houses)
	ans := 0
	n := len(heaters) - 1

	for i := 0; i < len(houses); i++ {
		left, right := 0, n
		for right-left > 1 {
			mid := left + (right-left)/2
			// 正好相等
			if houses[i] == heaters[mid] {
				left, right = mid, mid
				// 房子子加热站的右边
			} else if houses[i] > heaters[mid] {
				left = mid
			} else {
				right = mid
			}
		}
		min := Min(Abs(houses[i], heaters[left]), Abs(houses[i], heaters[right]))
		if ans < min {
			ans = min
		}
	}
	return ans
}

func Min(nums ...int) int {
	min := math.MaxInt64
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func Max(nums ...int) int {
	max := math.MinInt64
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func Abs(n1, n2 int) int {
	return int(math.Abs(float64(n1 - n2)))
}
