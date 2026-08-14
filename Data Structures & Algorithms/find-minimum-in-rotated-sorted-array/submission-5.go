func findMin(nums []int) int {
	minV := nums[0]
	for l, r := 0, len(nums)-1; l <= r; {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m - 1
		}
		if nums[m] < minV {
			minV = nums[m]
		}
	}
	return minV
}
