package middle

import "fmt"

func stringShift(s string, shift [][]int) string {
	d,m := 0,0
	for i:=0;i<len(shift);i++ {
		if d != shift[i][0] && m - shift[i][1] <= 0 {
			d = shift[i][0]
			m = shift[i][1] - m
		} else if d == shift[i][0] {
			d = shift[i][0]
			m = shift[i][1] + m
		} else {
			m = m -shift[i][1]
		}
	}

	if m % len(s) == 0 {
		return s
	}

	m %= len(s)

	if d == 0 {
		fmt.Println(s[m:],s[:m])
		return s[m:] + s[:m]
	}
	return s[len(s) - m:] + s[:len(s) - m]
}