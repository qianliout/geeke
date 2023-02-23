package unionfind

type UnionFind struct {
	Parent []int
	Rank   []int
	Count  int
}

func NewUnionFind(totalNodes int) *UnionFind {
	p := make([]int, totalNodes)
	r := make([]int, totalNodes)
	for i := 0; i < totalNodes; i++ {
		p[i] = i
		r[i] = 1
	}
	return &UnionFind{Parent: p, Rank: r, Count: totalNodes}
}

func (u *UnionFind) Find(x int) int {
	if u.Parent[x] != x {
		// 路径压缩
		u.Parent[x] = u.Find(u.Parent[x])
	}
	return u.Parent[x]
}

// 把x y 弄到一个集合中去
func (u *UnionFind) Union(x, y int) {
	xRoot := u.Find(x)
	yRoot := u.Find(y)

	// 把低的rank赋值给高的node，这里有些不理解，
	if xRoot != yRoot {
		u.Count--
		if u.Rank[xRoot] > u.Rank[yRoot] {
			u.Parent[yRoot] = xRoot
		} else if u.Rank[xRoot] < u.Rank[yRoot] {
			u.Parent[xRoot] = yRoot
		} else {
			u.Parent[yRoot] = xRoot
			u.Rank[xRoot]++
		}
	}
}

func (u *UnionFind) IsConnected(x, y int) bool {
	return u.Find(x) == u.Find(y)
}
