type point struct {
	r, c int
}

func countPaths(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	visited := make(map[point]bool, 0)
	return dfs(grid, 0, 0, visited)
}

func dfs(grid [][]int, r, c int, visited map[point]bool) int {
	if r < 0 || c < 0 {
		return 0
	}
	if r >= len(grid) || c >= len(grid[0]) {
		return 0
	}
	if visited[point{r, c}] {
		return 0
	}
	if grid[r][c] != 0 {
		return 0
	}
	if r == len(grid)-1 && c == len(grid[0])-1 {
		return 1
	}
	visited[point{r, c}] = true
	paths := 0
	paths += dfs(grid, r-1, c, visited)
	paths += dfs(grid, r, c+1, visited)
	paths += dfs(grid, r+1, c, visited)
	paths += dfs(grid, r, c-1, visited)
	delete(visited, point{r, c})
	return paths
}