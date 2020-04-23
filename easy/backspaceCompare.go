package easy

func backspaceCompare(S string, T string) bool {
	i:=len(S)-1
	j:=len(T)-1

	for i>0 || j > 0 {
		i = helper(S,i)
		j = helper(T,j)

		if i >0 && j>0 && S[i] != T[j] {
			return false
		}

		if (i>=0) != (j>=0) {
			return false
		}

		i--
		j--
	}

	return true
}

func helper(s string,pos int) int {
	skipCount := 0
	for pos >= 0 {
		if s[pos] == '#' {
			skipCount++
			pos--
		} else if skipCount > 0 {
			pos--
			skipCount--
		} else {
			break
		}
	}

	return pos
}