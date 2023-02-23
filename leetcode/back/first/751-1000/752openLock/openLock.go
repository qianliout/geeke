package main

import (
	"fmt"
)

func main() {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
	fmt.Println(openLock([]string{"0000"}, "8888"))
}

/*
你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。
每个拨轮可以自由旋转：例如把 '9' 变为  '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
字符串 target 代表可以解锁的数字，你需要给出最小的旋转次数，如果无论如何不能解锁，返回 -1。
示例 1:
输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
输出：6
解释：
可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
因为当拨动到 "0102" 时这个锁就会被锁定。
示例 2:
输入: deadends = ["8888"], target = "0009"
输出：1
解释：
把最后一位反向旋转一次即可 "0000" -> "0009"。
示例 3:
输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
输出：-1
解释：
无法旋转到目标数字且不被锁定。
示例 4:
输入: deadends = ["0000"], target = "8888"
输出：-1
提示：
死亡列表 deadends 的长度范围为 [1, 500]。
目标数字 target 不会在 deadends 之中。
每个 deadends 和 target 中的字符串的数字会在 10,000 个可能的情况 '0000' 到 '9999' 中产生。
*/
func openLock(deadends []string, target string) int {
	dead := make(map[string]bool)
	for _, d := range deadends {
		dead[d] = true
	}
	return bfs(target, dead)
}

func bfs(target string, used map[string]bool) int {
	res := 0
	queue := make([][]byte, 0)
	queue = append(queue, []byte("0000"))
	for len(queue) > 0 {
		ths := make([][]byte, 0)

		for j := 0; j < len(queue); j++ {
			first := queue[j]
			if used[string(first)] {
				continue
			}
			used[string(first)] = true // 在这里进行判断，是这道题的关键
			if string(first) == target {
				return res
			}
			for i := 0; i < len(first); i++ {
				after := append([]byte{}, first...)
				after[i] = add(after[i])
				if !used[string(after)] {
					ths = append(ths, after)
				}
			}
			for i := 0; i < len(first); i++ {
				after := append([]byte{}, first...)
				after[i] = sub(after[i])
				if !used[string(after)] {
					ths = append(ths, after)
				}
			}
		}
		queue = ths
		res++
	}
	return -1
}
func add(b byte) byte {
	if b == '9' {
		return '0'
	}
	return b + 1
}

func sub(b byte) byte {
	if b == '0' {
		return '9'
	}
	return b - 1
}
