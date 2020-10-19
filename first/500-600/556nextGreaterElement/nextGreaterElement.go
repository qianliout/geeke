package main

import (
	"strconv"
)

func main() {

}

/*
给定一个32位正整数 n，你需要找到最小的32位整数，其与 n 中存在的位数完全相同，并且其值大于n。如果不存在这样的32位整数，则返回-1。
示例 1:
输入: 12
输出: 21
示例 2:
输入: 21
输出: -1
*/

/*
1、从低位到高位，如果每位数字是依次递增的，那么就不存在符合题意的数，比如44332211。
2、需要从低位到高位遍历，判断是否能找到一个比高位某数更大的最小值。以12443111为例，
可以发现低位的3是比高位的2大的最小值(4也比2大，但不是最小值)，那么可以将3和2进行交换，得到13442111，
然后再把4以后的数按从小到大的顺序进行排列即可，答案为13111244.
3、怎样找到比低位大的最小值呢？可以考虑使用栈，从后往前，维持一个递增栈，当遇到小于栈顶的元素时(索引为i)，
依次出栈，并将栈顶元素的索引保存下来，从而找到低位比当前元素大的最小值所在的位置idx，将两个元素交换。然后从i+ 1开始，进行从小到大排序。即可得到答案。
*/
func nextGreaterElement(n int) int {
	nums := make([]int, 0)
	s := []byte(strconv.Itoa(n))
	for i := range s {
		nums = append(nums, int(s[i]-'0'))
	}
	stark := make([]int, 0)
	i := len(nums) - 1
	for i >= 0 {
		if len(stark) == 0 {
			stark = append(stark, i)
		} else {
			if nums[i] >= stark[len(stark)] {
				stark = append(stark, i)
			} else {
				break
			}
		}
		i--
	}
	if i == 0 {
		return -1
	}
	//

}
