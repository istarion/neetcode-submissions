type Graph struct {
	edges map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		edges: make(map[int][]int, 0),
	}
}

func (g *Graph) AddEdge(src, dst int) {
	if _, ok := g.edges[src]; !ok {
		g.edges[src] = []int{}
	}
	g.edges[src] = append(g.edges[src], dst)

}

func (g *Graph) RemoveEdge(src, dst int) bool {
	if _, ok := g.edges[src]; !ok {
		return false
	}
	for i, v := range g.edges[src] {
		if v == dst {
			g.edges[src] = append(g.edges[src][:i], g.edges[src][i+1:]...)
			return true
		}
	}
	return false
}

func (g *Graph) HasPath(src, dst int) bool {
	if src == dst {
		return true
	}
	visited := make(map[int]bool)
	return g.hasPathDFS(src, dst, visited)
}

func (g *Graph) hasPathDFS(src, dst int, visited map[int]bool) bool {
	if src == dst {
		return true
	}
	visited[src] = true
	if _, ok := g.edges[src]; !ok {
		return false
	}
	for _, neighbor := range g.edges[src] {
		if !visited[neighbor] {
			if g.hasPathDFS(neighbor, dst, visited) {
				return true
			}
			visited[neighbor] = true
		}
	}
	return false
}