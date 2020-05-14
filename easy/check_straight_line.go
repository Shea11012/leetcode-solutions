package easy

func checkStraightLine(coordinates [][]int) bool {
	// 此处使用数学上的斜率公式 a/b = c/d 交叉相乘得到 ad = cb
	y := coordinates[1][1] - coordinates[0][1]
	x := coordinates[1][0] - coordinates[0][0]
	for i := 2; i < len(coordinates); i++ {
		before := coordinates[i-1]
		current := coordinates[i]
		if y*(current[0]-before[0]) != x*(current[1]-before[1]) {
			return false
		}
	}
	return true
}
