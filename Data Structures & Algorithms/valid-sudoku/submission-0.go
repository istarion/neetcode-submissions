
func isValidSudoku(board [][]byte) bool {
	rowBitmaps := make([]int, 9)
	colBitmaps := make([]int, 9)
	boxBitmaps := make([]int, 9)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := board[i][j] - '0'
			if rowBitmaps[i]&(1<<num) != 0 {
				return false
			}
			if colBitmaps[j]&(1<<num) != 0 {
				return false
			}
			boxnum := (i / 3) + 3*(j/3)
			if boxBitmaps[boxnum]&(1<<num) != 0 {
				return false
			}
			rowBitmaps[i] |= 1 << num
			colBitmaps[j] |= 1 << num
			boxBitmaps[boxnum] |= 1 << num
		}
	}
	return true
}