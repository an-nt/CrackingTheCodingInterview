package codesignal

//Sudoku checks whether a 9x9 matrix is a valid sudoku puzzle or not
func Sudoku(grid [][]string) bool {
	for i := 0; i < len(grid); i += 3 {
		for j := 0; j < len(grid); j += 3 {
			sqrtResult := checkSqrt(grid, i, j)
			if !sqrtResult {
				return sqrtResult
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		rowResult := checkRow(grid[i])
		if !rowResult {
			return rowResult
		}
	}
	for i := 0; i < len(grid); i++ {
		colResult := checkCol(grid, i)
		if !colResult {
			return colResult
		}
	}
	return true
}

func checkSqrt(grid [][]string, row int, col int) bool {
	recorder := make(map[string]bool)
	for i := row; i < row+3; i++ {
		for j := col; j < col+3; j++ {
			if grid[i][j] != "." {
				_, existed := recorder[grid[i][j]]
				if existed {
					return false
				}
				recorder[grid[i][j]] = true
			}
		}
	}
	return true
}

func checkRow(row []string) bool {
	recorder := make(map[string]bool)
	for _, value := range row {
		if value != "." {
			_, existed := recorder[value]
			if existed {
				return false
			}
			recorder[value] = true
		}
	}
	return true
}

func checkCol(grid [][]string, colNum int) bool {
	recorder := make(map[string]bool)
	for _, row := range grid {
		if row[colNum] != "." {
			_, existed := recorder[row[colNum]]
			if existed {
				return false
			}
			recorder[row[colNum]] = true
		}
	}
	return true
}
