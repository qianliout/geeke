package unionfind

type UnionFind struct {
	Parent []int
	Rank   []int
}

func NewUnionFind(totalNodes int) *UnionFind {
	p := make([]int, totalNodes)
	r := make([]int, totalNodes)
	for i := 0; i < totalNodes; i++ {
		p[i] = i
		r[i] = 1
	}
	return &UnionFind{Parent: p, Rank: r}
}

func (u *UnionFind) Find(x int) int {
	if u.Parent[x] != x {
		u.Parent[x] = u.Find(u.Parent[x])
	}
	return u.Parent[x]
}

// 把x y 弄到一个集合中去
func (u *UnionFind) Union(x, y int) {
	xRoot := u.Find(x)
	yRoot := u.Find(y)
	if xRoot == yRoot {
		return
	}

	if xRoot != yRoot {
		if u.Rank[xRoot] > u.Rank[yRoot] {
			u.Parent[yRoot] = xRoot
		} else if u.Rank[xRoot] < u.Rank[yRoot] {
			u.Parent[xRoot] = yRoot
		} else {
			u.Parent[xRoot] = yRoot
			u.Rank[yRoot]++
		}
	}
}

func (u *UnionFind) IsConnected(x, y int) bool {
	return u.Find(x) == u.Find(y)
}
