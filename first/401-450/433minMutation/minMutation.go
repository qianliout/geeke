package main

import (
	"fmt"
)

func main() {
	start := "AAAAAAAA"
	end := "CCCCCCCC"
	bank := []string{"AAAAAAAA", "AAAAAAAC", "AAAAAACC", "AAAAACCC", "AAAACCCC", "AACACCCC", "ACCACCCC", "ACCCCCCC", "CCCCCCCA"}

	res := minMutation(start, end, bank)
	fmt.Println("res is ", res)
}

/*
一条基因序列由一个带有8个字符的字符串表示，其中每个字符都属于 "A", "C", "G", "T"中的任意一个。
假设我们要调查一个基因序列的变化。一次基因变化意味着这个基因序列中的一个字符发生了变化。
例如，基因序列由"AACCGGTT" 变化至 "AACCGGTA" 即发生了一次基因变化。
与此同时，每一次基因变化的结果，都需要是一个合法的基因串，即该结果属于一个基因库。
现在给定3个参数 — start, end, bank，分别代表起始基因序列，目标基因序列及基因库，请找出能够使起始基因序列变化为目标基因序列所需的最少变化次数。如果无法实现目标变化，请返回 -1。
注意:
    起始基因序列默认是合法的，但是它并不一定会出现在基因库中。
    所有的目标基因序列必须是合法的。
    假定起始基因序列与目标基因序列是不一样的。
示例 1:
start: "AACCGGTT"
end:   "AACCGGTA"
bank: ["AACCGGTA"]
返回值: 1
示例 2:
start: "AACCGGTT"
end:   "AAACGGTA"
bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]
返回值: 2
示例 3:
start: "AAAAACCC"
end:   "AACCCCCC"
bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]
返回值: 3
*/

func minMutation(start string, end string, bank []string) int {
	if start == end || start == "" || end == "" || len(bank) == 0 {
		return 0
	}
	bankMap := make(map[string]bool)
	for _, b := range bank {
		bankMap[b] = true
	}
	if i, ok := bankMap[end]; !ok || !i {
		return -1
	}
	queue := make([]Item, 0)
	queue = append(queue, Item{dna: start, count: 0})

	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		// 改
		for i, ch := range []byte(first.dna) {
			for _, b := range []byte{'A', 'C', 'G', 'T'} {
				if ch != b {
					newDNA := []byte(first.dna)
					newDNA[i] = b
					if string(newDNA) == end {
						return first.count + 1
					}
					if bankMap[string(newDNA)] {
						nitem := Item{dna: string(newDNA), count: first.count + 1}
						queue = append(queue, nitem)
						bankMap[nitem.dna] = false
					}

				}
			}
		}
	}
	return -1
}

type Item struct {
	dna   string
	count int
}
