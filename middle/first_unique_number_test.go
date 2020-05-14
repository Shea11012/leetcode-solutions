package middle

import (
	"fmt"
	"testing"
)

func Test_newListFer(t *testing.T) {
	l := newListFer()
	fmt.Printf("%p\n",l)
}

func Test_FirstUniqueNumber(t *testing.T) {
	cases := []struct{
		name string
		args struct{
			nums []int
		}
		want int
	}{
		{
			name:"one",
			args: struct{ nums []int }{nums: []int{2,3,4}},
			want:2,
		},
		{
			name:"two",
			args: struct{ nums []int }{nums: []int{2,3,4,2,3}},
			want:4,
		},
	}

	for _,v := range cases {
		t.Run(v.name,func (t *testing.T) {
			c := Constructor(v.args.nums)
			if r := c.ShowFirstUnique(); r != v.want {
				t.Errorf("实际输出%d,期望输出%d\n",r,v.want)
			}
		})
	}
}

func TestFirstUnique_Add(t *testing.T) {
	want := -1
	c := Constructor([]int{2,3,4})
	c.Add(2)
	c.Add(3)
	c.Add(4)
	c.Add(3)
	r := c.ShowFirstUnique()
	if r != want {
		t.Errorf("实际输出%d，期望输出%d\n",r,want)
	}
}