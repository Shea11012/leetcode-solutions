package middle

import "testing"

func Test_searchRotate(t *testing.T) {
	nums := []int{4,5,6,7,0,1,2}
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"one",
			args:args{
				nums:  nums ,
				target: 0,
			},
			want:4,
		},
		{
			name:"two",
			args:args{
				nums:   []int{5,1,3},
				target: 5,
			},
			want:0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRotate(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchRotate() = %v, want %v", got, tt.want)
			}
		})
	}
}