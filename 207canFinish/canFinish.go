package main

import (
	"fmt"
)

func main() {
	numCourses := 2
	prerequisites := [][]int{{1, 0}, {0, 3}}
	can := canFinish(numCourses, prerequisites)
	fmt.Println("can ", can)
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	directedGraph := make([]int, numCourses) // 入度，对应这表示这个课程他有多少个前置课程
	front := make(map[int][]int)             // map[key][]int 表示课程key,是哪些课程的前置课程

	queue := make([]int, 0) // 表示可以马上学习的课程，也就是入度为0的课程
	order := make([]int, 0)

	for _, cou := range prerequisites {
		if len(cou) != 2 {
			continue
		}
		directedGraph[cou[0]]++
		if front[cou[1]] == nil {
			front[cou[1]] = make([]int, 0)
		}
		front[cou[1]] = append(front[cou[1]], cou[0])
	}
	for i := range directedGraph {
		if directedGraph[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		order = append(order, first)
		for _, cu := range front[first] {
			directedGraph[cu]--
			if directedGraph[cu] == 0 {
				queue = append(queue, cu)
			}
		}
	}

	return len(order) == numCourses
}
