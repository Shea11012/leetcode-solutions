package easy

func buyAndSellStock(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if v := prices[i] - prices[i-1];v>0 {
			profit+=v
		}
	}

	return profit
}
