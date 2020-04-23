package middle

func searchRotate(nums []int,target int) int {
	left,right := 0,len(nums)-1

	for left < right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		/*if nums[left] <= target && target <= nums[mid] {
			right = mid
		} else if target >= nums[mid] && nums[mid] < nums[left] {
			right = mid
		} else {
			left = mid + 1
		}*/
		if nums[left] <= nums[mid] && target <= nums[mid] && nums[left] <= target { // 右边界左移  第一种情况
			right = mid
		} else if nums[left] > nums[mid] && (target >= nums[left] || target <= nums[mid]) { // 右边界左移  第二种情况
			right = mid;
		} else { // 其余情况左边界右移
			left = mid + 1
		}
	}

	if left == right && nums[left] == target {
		return left
	}

	return -1
}
