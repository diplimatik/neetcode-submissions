func isValidSudoku(board [][]byte) bool {
	var mapArr [3][3]map[byte]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			mapArr[i][j] = make(map[byte]int)
		}
	}
	for r := 0; r < 9; r++ {
		cMap := make(map[byte]struct{}, 9)
		rMap := make(map[byte]struct{}, 9)
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' {
				if _, ok := cMap[board[r][c]]; ok {
					return false
				} else if val := mapArr[r/3][c/3][board[r][c]]; val > 2 {
					return false
				}
				cMap[board[r][c]] = struct{}{}
				mapArr[r/3][c/3][board[r][c]]++
			}
			if board[c][r] != '.' {
				if _, ok := rMap[board[c][r]]; ok {
					return false
				} else if val := mapArr[c/3][r/3][board[c][r]]; val > 2 {
					return false
				}
				rMap[board[c][r]] = struct{}{}
				mapArr[c/3][r/3][board[c][r]]++
			}
		}
	}
	
	return true
}