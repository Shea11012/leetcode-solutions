package easy

import "testing"

func Test_LastStoneWeight(t *testing.T) {
	cases := map[int][]int{
		1:[]int{3,2,6,1,8,10,11},
	}
	for want,i	:= range cases {
		if r := lastStoneWeight(i);r != want {
			t.Errorf("want:%d,error:%d",want,r)
		}
	}
}
