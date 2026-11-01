type PathToVertex struct {
	p int
	v int
}
type PathToVertexHeap []PathToVertex

func (h PathToVertexHeap) Len() int           { return len(h) }
func (h PathToVertexHeap) Less(i, j int) bool { return h[i].p < h[j].p }
func (h PathToVertexHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PathToVertexHeap) Push(x any) {
	*h = append(*h, x.(PathToVertex))
}
func (h *PathToVertexHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = PathToVertex{}
	*h = old[:n-1]
	return item
}

type Edge struct {
	v, w int
}

func shortestPath(n int, edges [][]int, src int) map[int]int {
	h := make(PathToVertexHeap, 0)
	heap.Init(&h)
	result := make(map[int]int, 0)
	heap.Push(&h, PathToVertex{p: 0, v: src})
	edgemap := make(map[int][]Edge, 0)
	for _, e := range edges {
		if _, ok := edgemap[e[0]]; !ok {
			edgemap[e[0]] = make([]Edge, 0)
		}
		edgemap[e[0]] = append(edgemap[e[0]], Edge{e[1], e[2]})
	}

	for h.Len() > 0 {
		ptv := heap.Pop(&h).(PathToVertex)
		if _, ok := result[ptv.v]; ok {
			continue
		}
		result[ptv.v] = ptv.p
		if len(result) == n {
			break
		}
		for _, neighbour := range edgemap[ptv.v] {
			heap.Push(&h, PathToVertex{p: ptv.p + neighbour.w, v: neighbour.v})
		}
	}
	for i := 0; i < n; i++ {
		if _, ok := result[i]; !ok {
			result[i] = -1
		}
	}
	return result
}
