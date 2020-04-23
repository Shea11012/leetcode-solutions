package easy

// 分治法，将问题分解为子问题递归解决
func divide(nums []int, left, right int) int {
	return 0
}

// 动态规划
func maxSubArray2(nums []int) int {
	ans, sum := nums[0], 0
	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}
		if ans < sum {
			ans = sum
		}
	}
	return ans
}
