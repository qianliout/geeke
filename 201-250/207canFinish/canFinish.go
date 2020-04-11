package main

import (
	"fmt"
)

func main() {
	prerequisites := [][]int{{1, 0}}

	res := canFinish(2, prerequisites)
	fmt.Println("can finish is ", res)
}

/*
你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。
在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？
示例 1:
输入: 2, [[1,0]]
输出: true
解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
示例 2:
输入: 2, [[1,0],[0,1]]
输出: false
解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。
提示：
    输入的先决条件是由 边缘列表 表示的图形，而不是 邻接矩阵 。详情请参见图的表示法。
    你可以假定输入的先决条件中没有重复的边。
    1 <= numCourses <= 10^5
链接：https://leetcode-cn.com/problems/course-schedule
*/
// bfs 容易理解
func canFinish(numCourses int, prerequisites [][]int) bool {
	indegrees := make([]int, numCourses)
	// 注意理解这个变量，这里这个变量是最重要的，他的意思是，当前课程，有多少个对应的前置课程
	adjacency := make([][]int, numCourses)
	// 先把所有的课程加进去,注意，这里是当前课程指向前置课程，所以是长前置课程的入度
	for _, req := range prerequisites {
		indegrees[req[0]]++
		adjacency[req[1]] = append(adjacency[req[1]], req[0]) // 计算出，当前课程有多少个前置课程
	}

	queue := make([]int, 0)
	// 先把初始条件下，出度为0的课程找出来，放到队列中
	for i := 0; i < numCourses; i++ {
		if indegrees[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		pre := queue[0]
		numCourses--
		queue = queue[1:len(queue)]
		for _, cu := range adjacency[pre] {
			indegrees[cu]--
			if indegrees[cu] == 0 {
				queue = append(queue, cu)
			}
		}
	}
	return numCourses == 0
}

func canFinish2(numCourses int, prerequisites [][]int) bool {
	flag := make([]int, numCourses)
	adjacency := make([][]int, numCourses)
	for _, req := range prerequisites {
		adjacency[req[1]] = append(adjacency[req[1]], req[0]) // 计算出，当前课程有多少个前置课程
	}
	for i := 0; i < numCourses; i++ {
		if !dfs(adjacency, flag, i) {
			return false
		}
	}
	return true

}
func dfs(adjacency [][]int, flag []int, i int) bool {
	if flag[i] == 1 {
		return false
	}
	if flag[i] == -1 {
		return true
	}
	flag[i] = 1
	for _, curr := range adjacency[i] {
		if !dfs(adjacency, flag, curr) {
			return false
		}
	}
	flag[i] = -1
	return true
}
