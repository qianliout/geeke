package main

func main() {

}

func findMinHeightTrees(n int, edges [][]int) []int {
	inDegree := make(map[int][]int)
	for i := 0; i < n; i++ {
		inDegree[i] = make([]int, 0)
	}
	// 计算入度
	for _, e := range edges {
		inDegree[e[1]] = append(inDegree[e[1]], e[0])
		inDegree[e[0]] = append(inDegree[e[0]], e[1])
	}
	// 循环删除入度为1的节点
	for len(inDegree) > 2 {
		queue := make([][]int, 0)
		for k, v := range inDegree {
			if len(v) == 1 {
				queue = append(queue, []int{k, v[0]})
			}
		}
		for _, q := range queue {
			first := q[0]
			second := q[1]

			delete(inDegree, first)

			gre := inDegree[second]
			for i, k := range gre {
				if k == first {
					s := append(gre[:i], gre[i+1:]...)
					inDegree[second] = s
				}
			}
		}
	}

	res := make([]int, 0)
	for k := range inDegree {
		res = append(res, k)
	}
	return res
}
