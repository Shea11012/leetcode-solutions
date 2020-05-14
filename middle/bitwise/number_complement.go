package bitwise

import "math"

func findComplement(num int) int {
	var count float64
	ans := 0
	for num > 0 {
		t := num % 2
		num /= 2
		if t == 0 {
			ans += 1 * int(math.Pow(2,count))
		}
		count++
	}
	return ans
}
