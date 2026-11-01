type point struct {
	r, c int
}

func (p point) plus(other point) point {
	return point{p.r + other.r, p.c + other.c}
}

func shortestPath(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	ROWS, COLS := len(grid), len(grid[0])
	if COLS == 0 {
		return 0
	}
	if grid[0][0] != 0 {
		return -1
	}
	queue := make([]point, 0)
	queue = append(queue, point{0, 0})
	visited := make(map[point]bool)
	visited[point{0, 0}] = true

	sp := 0
	for len(queue) > 0 {
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			p := queue[i]
			if p.r == ROWS-1 && p.c == COLS-1 {
				return sp
			}
			neighbours := []point{
				point{-1, 0}, point{0, 1}, point{1, 0}, point{0, -1},
			}
			for _, dp := range neighbours {
				nextPoint := p.plus(dp)
				if nextPoint.c < 0 || nextPoint.r < 0 ||
					nextPoint.c == COLS || nextPoint.r == ROWS ||
					visited[nextPoint] || grid[nextPoint.r][nextPoint.c] != 0 {
					continue
				}
				queue = append(queue, nextPoint)
				visited[nextPoint] = true
			}
		}
		sp++
		queue = queue[qLen:]
	}
	return -1
}