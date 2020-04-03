package easy

import "testing"

func Test_MaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	if v := maxSubArray(nums); v != 6 {
		t.Errorf("error :%d", v)
	}
}
