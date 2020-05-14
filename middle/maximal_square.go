package middle

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	matrixH,matrixL := len(matrix),len(matrix[0])
	dp := make([]int,matrixL+1)
	ans := 0
	left := 0
	for i := 0; i < matrixH; i++ {
		for j := 1; j <= matrixL; j++ {
			t := dp[j] // 1
			if matrix[i][j-1] == '1' {
				// minS 内的 dp[j] 表示二维数组的上方，dp[j-1]表示左边，left表示左上方
				dp[j] = minS(dp[j-1],left,dp[j]) + 1
				ans = maxS(ans,dp[j])
			} else {
				dp[j] = 0 // 需要手动将 dp[j] 置为 0，dp 是一维数组会保留上一行的值
			}
			left = t // 2 数字标记的地方，将 j-1 的值保存到了 left 中
		}
	}
	return ans * ans
}

// 空间复杂度 m*n
func dp1(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	matrixH, matrixL := len(matrix), len(matrix[0])
	dp := make([][]int, matrixH+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, matrixL+1)
	}
	ans := 0
	for i := 1; i <= matrixH; i++ {
		for j := 1; j <= matrixL; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = minS(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
				ans = maxS(ans, dp[i][j])
			}
		}
	}
	return ans * ans
}

// 空间复杂度 n


func maxS(x,y int) int {
	if x > y {
		return x
	}

	return y
}

func minS(x... int) int {
	m := x[0]
	for i:=1; i<len(x);i++ {
		if x[i] < m {
			m = x[i]
		}
	}

	return m
}
