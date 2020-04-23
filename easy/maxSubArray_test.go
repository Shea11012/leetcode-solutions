package easy

import "testing"

func Test_MaxSubArray(t *testing.T) {
	want := 6
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	if v := maxSubArray2(nums); v != want {
		t.Errorf("error :%d", v)
	}
}
