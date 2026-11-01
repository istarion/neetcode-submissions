func countPaths(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	return dfs(grid, 0, 0)
}

func dfs(grid [][]int, r, c int) int {
	if r < 0 || c < 0 {
		return 0
	}
	if r >= len(grid) || c >= len(grid[0]) {
		return 0
	}
	if grid[r][c] != 0 {
		return 0
	}
	if r == len(grid)-1 && c == len(grid[0])-1 {
		return 1
	}
	grid[r][c] = 1
	paths := 0
	paths += dfs(grid, r-1, c)
	paths += dfs(grid, r, c+1)
	paths += dfs(grid, r+1, c)
	paths += dfs(grid, r, c-1)
	grid[r][c] = 0
	return paths
}