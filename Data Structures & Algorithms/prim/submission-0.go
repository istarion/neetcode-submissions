type PathToEdge struct {
	p int
	e Edge
}
type PathToEdgeHeap []PathToEdge

func (h PathToEdgeHeap) Len() int           { return len(h) }
func (h PathToEdgeHeap) Less(i, j int) bool { return h[i].p < h[j].p }
func (h PathToEdgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PathToEdgeHeap) Push(x any) {
	*h = append(*h, x.(PathToEdge))
}
func (h *PathToEdgeHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = PathToEdge{}
	*h = old[:n-1]
	return item
}

type Edge struct {
	src, dst int
}

type WeightedEdge struct {
	dst, w int
}

func minimumSpanningTree(n int, edges [][]int) int {
	h := make(PathToEdgeHeap, 0)
	heap.Init(&h)
	edgemap := make(map[int][]WeightedEdge, 0)
	for _, e := range edges {
		if _, ok := edgemap[e[0]]; !ok {
			edgemap[e[0]] = make([]WeightedEdge, 0)
		}
		if _, ok := edgemap[e[1]]; !ok {
			edgemap[e[1]] = make([]WeightedEdge, 0)
		}
		edgemap[e[0]] = append(edgemap[e[0]], WeightedEdge{e[1], e[2]})
		edgemap[e[1]] = append(edgemap[e[1]], WeightedEdge{e[0], e[2]})
	}
	visited := make(map[int]any, 0)
	visited[0] = struct{}{}
	for _, neighbour := range edgemap[0] {
		heap.Push(&h, PathToEdge{
			p: neighbour.w,
			e: Edge{src: 0, dst: neighbour.dst},
		})
	}
	result := 0

	for h.Len() > 0 {
		pte := heap.Pop(&h).(PathToEdge)
		if _, ok := visited[pte.e.dst]; ok {
			continue
		}
		visited[pte.e.dst] = struct{}{}
		result += pte.p
		if len(visited) == n {
			break
		}
		for _, neighbour := range edgemap[pte.e.dst] {
			if _, ok := visited[neighbour.dst]; ok {
				continue
			}
			heap.Push(&h, PathToEdge{
				p: neighbour.w,
				e: Edge{pte.e.dst, neighbour.dst},
			})
		}
	}
	if len(visited) < n {
		return -1
	}
	return result
}
