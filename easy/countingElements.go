package easy

func countElements(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	arrMap := make(map[int]int)
	ans := 0
	for _, v := range arr {
		arrMap[v] = 0
	}
	for _,v := range arr {
		r := v+1
		if _, ok := arrMap[r]; ok {
			arrMap[r]++
		}
	}

	for i := range arrMap {
		ans += arrMap[i]
	}

	return ans
}
