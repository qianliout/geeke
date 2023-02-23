package main

import (
	"fmt"
)

func main() {
	res := canMeasureWater(3, 5, 4)
	fmt.Println("res is ", res)
}

/*
有两个容量分别为 x升 和 y升 的水壶以及无限多的水。请判断能否通过使用这两个水壶，从而可以得到恰好 z升 的水？
如果可以，最后请用以上水壶中的一或两个来盛放取得的 z升 水。
你允许：
    装满任意一个水壶
    清空任意一个水壶
    从一个水壶向另外一个水壶倒水，直到装满或者倒空
示例 1: (From the famous "Die Hard" example)
输入: x = 3, y = 5, z = 4
输出: True
示例 2:
输入: x = 2, y = 6, z = 5
输出: False
*/

// 这道题的难点是定义状态，和最后结果时返回
func canMeasureWater(x int, y int, z int) bool {
	if z == 0 {
		return true
	}
	if x+y < z {
		return false
	}

	// [2]int  表示AB两个桶的水量,这是这道题的难点
	queue := [][2]int{{0, 0}}
	visited := make(map[[2]int]bool)
	for len(queue) > 0 {
		head := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		curX := head[0]
		curY := head[1]
		if curX+curY == z || curX == z || curY == z {
			return true
		}
		nextStates := getNextStates(curX, curY, x, y)
		// fmt.Println("head to nextStates", head, nextStates)

		for _, s := range nextStates {
			if !visited[s] {
				queue = append(queue, s)
				visited[s] = true
			}
		}
	}
	return false
}

func getNextStates(curX, curY, x, y int) [][2]int {
	nextStatus := make([][2]int, 0)

	// 以下两个状态，对应操作 1
	// 外部加水，使A加满
	nextStatus1 := [2]int{x, curY}
	if curX < x { // 没有满的时候才加水
		nextStatus = append(nextStatus, nextStatus1)
	}

	// 外部加水，使用B加满
	nextStatus2 := [2]int{curX, y}
	if curY < y { // 没有满的时候才加水
		nextStatus = append(nextStatus, nextStatus2)
	}
	// 把A清空
	nextStatus3 := [2]int{0, curY}
	// 把B清空
	nextStatus4 := [2]int{curX, 0}

	// 有水才能清空
	if curX > 0 {
		nextStatus = append(nextStatus, nextStatus3)
	}
	if curY > 0 {
		nextStatus = append(nextStatus, nextStatus4)
	}

	// 从 A 到 B，使得 B 满，A 还有剩
	nextStatus5 := [2]int{curX - (y - curY), y}
	if curX > 0 && curX-(y-curY) > 0 {
		nextStatus = append(nextStatus, nextStatus5)
	}
	// 从 A 到 B，此时 A 的水太少，A 倒尽，B 没有满
	nextStatus6 := [2]int{0, curX + curY}
	if curX > 0 && curX < y-curY {
		nextStatus = append(nextStatus, nextStatus6)
	}

	// 从 B 到 A，使得 A 满，B 还有剩余
	nextState7 := [2]int{x, curY - (x - curX)}
	if curY > 0 && curY-(x-curX) > 0 {
		nextStatus = append(nextStatus, nextState7)
	}

	// 从 B 到 A，此时 B 的水太少，B 倒尽，A 没有满
	nextState8 := [2]int{curX + curY, 0}
	if curY > 0 && curY < x-curX {
		nextStatus = append(nextStatus, nextState8)
	}

	return nextStatus
}
