package middle

import "testing"

func Test_RotateMatrix(t *testing.T) {
	testcase := map[int][][]int{
		0:[][]int{
			{1,2,3},
			{4,5,6},
			{7,8,9},
		},
		1:[][]int{
			{5,1,9,11},
			{2,4,8,10},
			{13,3,6,7},
			{15,14,12,16},
		},
	}

	for _, v := range testcase {
		rotate(v)
	}
}
