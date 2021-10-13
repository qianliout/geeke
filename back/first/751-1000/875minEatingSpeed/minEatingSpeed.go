package main

import (
	"fmt"

	"outback/leetcode/back/common"
)

func main() {
	// nums := []int{873375536, 395271806, 617254718, 970525912, 634754347, 824202576, 694181619, 20191396, 886462834, 442389139, 572655464, 438946009, 791566709, 776244944, 694340852, 419438893, 784015530, 588954527, 282060288, 269101141, 499386849, 846936808, 92389214, 385055341, 56742915, 803341674, 837907634, 728867715, 20958651, 167651719, 345626668, 701905050, 932332403, 572486583, 603363649, 967330688, 484233747, 859566856, 446838995, 375409782, 220949961, 72860128, 998899684, 615754807, 383344277, 36322529, 154308670, 335291837, 927055440, 28020467, 558059248, 999492426, 991026255, 30205761, 884639109, 61689648, 742973721, 395173120, 38459914, 705636911, 30019578, 968014413, 126489328, 738983100, 793184186, 871576545, 768870427, 955396670, 328003949, 786890382, 450361695, 994581348, 158169007, 309034664, 388541713, 142633427, 390169457, 161995664, 906356894, 379954831, 448138536}

	nums := []int{312884470}

	speed := bin(nums, 968709470)
	fmt.Println(possible(nums, 5, 29))
	fmt.Println("res is ", speed, len(nums))
}

/*
珂珂喜欢吃香蕉。这里有 N 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 H 小时后回来。
珂珂可以决定她吃香蕉的速度 K （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 K 根。如果这堆香蕉少于 K 根，
她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。
珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。
返回她可以在 H 小时内吃掉所有香蕉的最小速度 K（K 为整数）。
示例 1：
输入: piles = [3,6,7,11], H = 8
输出: 4
示例 2：
输入: piles = [30,11,23,4,20], H = 5
输出: 30
示例 3：
输入: piles = [30,11,23,4,20], H = 6
输出: 23
提示：
1 <= piles.length <= 10^4
piles.length <= H <= 10^9
1 <= piles[i] <= 10^9
*/

//  二分法
func minEatingSpeed(piles []int, H int) int {
	if len(piles) <= 0 {
		return 0
	}
	maxPile := piles[0]
	for _, v := range piles {
		if v > maxPile {
			maxPile = v
		}
	}
	// fmt.Println("max pile", maxPile)
	left, right := 1, maxPile
	// 二分找区间,相当于找左区间(最小的左值)
	for left < right {
		mid := left + (right-left+1)/2
		p := possible(piles, H, mid)
		if p {
			//  下一轮的搜索区间是: [left,mid-1]
			right = mid - 1
		} else {
			left = mid
		}
	}

	// 判断最后一个(一定要做最后的判断)
	if possible(piles, H, left) {
		return left
	}

	return left + 1

}

func possible(piles []int, H, k int) bool {
	sum := 0
	for _, p := range piles {
		// 我们可以推断出珂珂将在 Math.ceil(p / K) = ((p-1) // K) + 1 小时内吃完这一堆，
		// 我们将每一堆的完成时间加在一起并与 H 比较 。
		// 向下取整后加1,防止是整数倍的情况(是整数倍会有什么问题呢? 加入是11个香蕉,如果每次吃5个,那就是3次,11/5+1,但是如果是10个香蕉,还这么算的法,就出错,所以)

		// 以下这种写法好理解
		// sum += p / k
		// if p%k != 0 {
		//	sum += 1
		// }

		sum += (p-1)/k + 1
	}
	return sum <= H
}

func bin(piles []int, H int) int {
	max := 0
	for _, p := range piles {
		max = common.Max(max, p)
	}
	// 开始二分
	left, right := 2, max
	for left <= right {
		mid := left + (right-left)/2
		if possible(piles, H, mid) {
			// 锁定右边界
			right = mid - 1
		} else if !possible(piles, H, mid) {
			left = mid + 1
		}
	}
	return left
}
