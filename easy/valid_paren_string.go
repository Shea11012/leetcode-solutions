package easy

/*
	贪心算法,匹配有效括号,如果是正确匹配总是会在一个正向的范围之内，如：[0,3]，[0,1]
 */
func checkValidString(s string) bool {
	lo,hi := 0,0
	for _, i := range s {
		if i == '(' {
			lo++
		} else if lo > 0 {
			lo--
		}

		if i != ')' {
			hi++
		} else {
			if hi <= 0 {
				return false
			}
			hi--
		}
	}

	return lo == 0
}
