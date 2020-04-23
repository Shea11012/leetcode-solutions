package middle

import "math"

func minPathSum(grid [][]int) int {
	//ans := dfsSum(grid, 0, 0)
	ans := dpMinPathNum(grid)
	return ans
}

func dfsSum(grid [][]int, x int, y int) int {
	if x == len(grid) || y == len(grid) {
		return math.MaxInt32
	}

	if x == len(grid)-1 && y == len(grid[0])-1 {
		return grid[x][y]
	}

	return grid[x][y] + min(dfsSum(grid, x, y+1), dfsSum(grid, x+1, y))
}

func dpMinPathNum(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j == 0 {
				grid[i][j] = grid[i][j]
			} else if i > 0 && j > 0 {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			} else if i == 0 && j > 0 {
				grid[i][j] += grid[i][j-1]
			} else if i > 0 && j == 0 {
				grid[i][j] += grid[i-1][j]
			}
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}

func min(x, y int) int {
	if x >= y {
		return y
	}

	return x
}
