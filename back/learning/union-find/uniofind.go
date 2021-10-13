package unionfind

type UnionFine struct {
	Count  int
	Parent []int
	Rank   []int
}

func NewUnionFind(grid [][]int) UnionFine {
	m, n := len(grid), len(grid[0])
	parent := make([]int, 0)
	rank := make([]int, 0)
	count := 0
	for i := 0; i < m*n; i++ {
		parent = append(parent, -1)
		rank = append(rank, 0)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				parent[i*n+j] = i*n + j
				count++
			}
		}
	}
	return UnionFine{
		Count:  count,
		Parent: parent,
		Rank:   rank,
	}
}

func (this *UnionFine) Find(i int) int {
	if this.Parent[i] != i {
		this.Parent[i] = this.Find(this.Parent[i])
	}
	return this.Parent[i]
}

func (this *UnionFine) Union(x, y int) {

	rootY := this.Find(y)
	rootX := this.Find(x)

	if rootY != rootX {
		if this.Rank[rootX] > this.Rank[rootY] {
			this.Parent[rootY] = rootX
		} else if this.Rank[rootX] < this.Rank[rootY] {
			this.Parent[rootX] = rootY
		} else {
			this.Parent[rootY] = rootX
			this.Rank[rootX]++
		}
		this.Count--
	}
}
