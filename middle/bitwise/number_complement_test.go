package bitwise

import "testing"

func Test_findComplement(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"one",
			args:args{num:10},
			want:5,
		},
		{
			name:"two",
			args:args{num:1},
			want:0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findComplement(tt.args.num); got != tt.want {
				t.Errorf("findComplement() = %v, want %v", got, tt.want)
			}
		})
	}
}