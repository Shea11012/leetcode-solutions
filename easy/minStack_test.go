package easy

import "testing"

func Test_Push(t *testing.T) {
	s := constructor()
	s.push(1)
	s.push(2)
	s.push(3)
	s.push(1)
}
