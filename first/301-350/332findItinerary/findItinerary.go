package main

import (
	"fmt"
	"sort"
)

func main() {
	// tickets := [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}
	tickets := [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}
	// tickets := [][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}}
	// tickets := [][]string{{"EZE", "AXA"}, {"TIA", "ANU"}, {"ANU", "JFK"}, {"JFK", "ANU"}, {"ANU", "EZE"}, {"TIA", "ANU"}, {"AXA", "TIA"}, {"TIA", "JFK"}, {"ANU", "TIA"}, {"JFK", "TIA"}}
	fmt.Println(len(tickets))
	res := findItinerary(tickets)
	fmt.Println("res is ", res)
}

/*
给定一个机票的字符串二维数组 [from, to]，子数组中的两个成员分别表示飞机出发和降落的机场地点，对该行程进行重新规划排序。
所有这些机票都属于一个从JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 出发。
说明:
    如果存在多种有效的行程，你可以按字符自然排序返回最小的行程组合。例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"]
相比就更小，排序更靠前
    所有的机场都用三个大写字母表示（机场代码）。
    假定所有机票至少存在一种合理的行程。
示例 1:
输入: [["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
输出: ["JFK", "MUC", "LHR", "SFO", "SJC"]
示例 2:
输入: [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
输出: ["JFK","ATL","JFK","SFO","ATL","SFO"]
解释: 另一种有效的行程是 ["JFK","SFO","ATL","JFK","ATL","SFO"]。但是它自然排序更大更靠后。
*/

func findItinerary(tickets [][]string) []string {
	// 因为有同样的机票，所以只能用个数判断
	pre := make(map[string]int, 0)
	for _, t := range tickets {
		key := fmt.Sprintf("%s_%s", t[0], t[1])
		pre[key]++
	}
	used := make(map[string]int)
	res := make([]string, 0)
	res = append(res, "JFK")
	path := make([]string, 0)
	var end bool
	dfs(tickets, pre, used, path, &res, "JFK", &end)
	return res
}

func dfs(tickets [][]string, pre, used map[string]int, path []string, res *[]string, last string, end *bool) {
	if *end {
		return
	}
	if len(path) == len(tickets) {
		*res = append(*res, append([]string{}, path...)...)
		*end = true
		return
	}

	newPath := make([]string, 0)
	for _, t := range tickets {
		if t[0] == last {
			newPath = append(newPath, t[1])
		}
	}

	sort.Sort(Item(newPath))
	for _, p := range newPath {
		key := fmt.Sprintf("%s_%s", last, p)
		if used[key] < pre[key] {
			used[key]++
			path = append(path, p)

			dfs(tickets, pre, used, path, res, p, end)

			used[key]--
			path = path[:len(path)-1]
		}
	}
}

type Item []string

func (it Item) Len() int {
	return len(it)
}

func (it Item) Less(i, j int) bool {
	if it[i] < it[j] {
		return true
	}
	return false
}

func (it Item) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}
