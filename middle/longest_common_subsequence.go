package middle

func longestCommonSubsequence(text1 string, text2 string) int {
	t1, t2 := len(text1), len(text2)
	ans := 0
	for i := 1; i <= t1; i++ {
		top := 0
		for j := 1; j <= t2; j++ {
			left := 0
			if text1[i-1] == text2[j-1] {
				ans++
				left = ans
			} else {
				left = max(left, top)
			}
			top = left
		}
	}

	return ans
}

func lcs1(text1 string, text2 string) int {
	t1, t2 := len(text1), len(text2)
	var dp = make([][]int, t1)
	for i := range dp {
		dp[i] = make([]int, t2)
	}

	for i := 0; i <= t1; i++ {
		for j := 0; j <= t2; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[t1][t2]
}

func max(m, n int) int {
	if m > n {
		return m
	}

	return n
}
