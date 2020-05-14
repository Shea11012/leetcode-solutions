package bitwise

func rangeBitWiseNumber(a,b int) int {
	return one(a,b)
}

// 通过取得 a 和 b 的最高有效位来判断
func one(a,b int) int {
	msb := func (x int) int {
		count := -1
		for x != 0 {
			x >>= 1
			count++
		}
		return count
	}
	ans := 0

	for (a & b) != 0 {
		msbA := msb(a)
		msbB := msb(b)

		// 如果最高有效位不同则直接退出
		if msbA != msbB {
			break
		}
		v := 1 << msbA
		ans += v	// 加上最高有效位的值
		// 去除最高有效位
		a -= v
		b -= v
	}

	return ans
}

// 通过减少 b 的 1
func two(a,b int) int {
	reduce1 := func(a,b int) int {
		for a < b {
			// 先计算 b & -b 取得值肯定是 1，然后减去 1，相当于减去 b 上有效 bit 位（1）
			b -= b & -b
		}
		return b
	}

	return a & reduce1(a,b)
}
