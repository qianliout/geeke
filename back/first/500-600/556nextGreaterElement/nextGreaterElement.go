package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(nextGreaterElement(1999999999))
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
	ss := []byte(strconv.Itoa(n))
	stark := make([]int, 0)
	// 找到要交换的第一个
	i, idx := len(ss)-1, 0
	for i >= 0 {
		for len(stark) > 0 && ss[i] < ss[stark[len(stark)-1]] {
			idx = stark[len(stark)-1]
			stark = stark[:len(stark)-1]
		}
		// 如果前面有比后面小的数，就交换了
		if idx > 0 {
			ss[i], ss[idx] = ss[idx], ss[i]
			break
		}
		stark = append(stark, i)
		i--
	}
	// 再排序后面的
	if idx == 0 {
		return -1
	}
	// fmt.Println("ss ", string(ss), i, idx)
	sort.Sort(item(ss[i+1:]))
	// fmt.Println("ss ", string(ss), i, idx)
	n, err := strconv.Atoi(string(ss))
	if err != nil || n > math.MaxInt32 {
		return -1
	}
	return n
}

type item []byte

func (it item) Len() int {
	return len(it)

}

func (it item) Less(i, j int) bool {
	return it[i] < it[j]
}

func (it item) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}
