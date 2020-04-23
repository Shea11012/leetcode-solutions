package middle

import "testing"

func Test_numOfIslands(t *testing.T) {
	tests := testHelper()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfIslands(tt.args.grid); got != tt.want {
				t.Errorf("numOfIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}

type args struct {
	grid [][]byte
}

func testHelper() []struct {
	name string
	args args
	want int
} {
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{grid: [][]byte{
				{1, 1, 1, 1, 0},
				{1, 1, 0, 1, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0},
			}},
			want: 1,
		},
		{
			name: "two",
			args: args{grid: [][]byte{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 0, 1, 1},
			}},
			want: 3,
		},
		{
			name: "three",
			args: args{grid: [][]byte{
				{1, 1, 1, 1, 0},
				{1, 1, 0, 1, 0},
				{1, 0, 1, 1, 0},
				{0, 0, 0, 1, 0},
			}},
			want: 1,
		},
	}
	return tests
}

func Test_numOfIslands2(t *testing.T) {
	tests := testHelper()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfIslands2(tt.args.grid); got != tt.want {
				t.Errorf("numOfIslands2() = %v, want %v", got, tt.want)
			}
		})
	}
}