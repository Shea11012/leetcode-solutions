package easy

import "testing"

func Test_checkValidString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name:"()",
			args:args{s:"()"},
			want:true,
		},
		{
			name:"(*)",
			args:args{s:"(*)"},
			want:true,
		},
		{
			name:"(*))",
			args:args{s:"(*))"},
			want:true,
		},
		{
			name:")(",
			args:args{s:")("},
			want:false,
		},
		{
			name:"()()",
			args:args{s:"()()"},
			want:true,
		},
		{
			name:"(*()",
			args:args{s:"(*()"},
			want:true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValidString(tt.args.s); got != tt.want {
				t.Errorf("checkValidString() = %v, want %v", got, tt.want)
			}
		})
	}
}