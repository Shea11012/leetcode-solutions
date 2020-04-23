package easy

import "testing"

func Test_BuyAndSellStock(t *testing.T) {
	tests := map[int][]int{
		7: []int{7, 1, 5, 3, 6, 4},
		4: []int{1, 2, 3, 4, 5},
		0: []int{7, 6, 4, 3, 1},
	}

	for want, prices := range tests {
		if v := buyAndSellStock(prices); v != want {
			t.Errorf("期望利润：%d，实际利润：%d", want, v)
		}
	}
}
