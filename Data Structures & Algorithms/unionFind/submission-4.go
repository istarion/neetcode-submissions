
type UnionFind struct {
	parents        []int
	size           []int
	num_components int
}

func NewUnionFind(n int) *UnionFind {
	parents := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
		size[i] = 1
	}
	return &UnionFind{
		parents:        parents,
		size:           size,
		num_components: n,
	}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parents[x] != uf.parents[uf.parents[x]] {
		uf.parents[x] = uf.Find(uf.parents[x])
	}
	return uf.parents[x]
}

func (uf *UnionFind) IsSameComponent(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX == rootY {
		return false
	}
	uf.num_components--
	if uf.size[rootX] < uf.size[rootY] {
		uf.parents[rootX] = rootY
		uf.size[rootY] += uf.size[rootX]
	} else {
		uf.parents[rootY] = rootX
		uf.size[rootX] += uf.size[rootY]
	}
	return true
}

func (uf *UnionFind) GetNumComponents() int {
	return uf.num_components
}
