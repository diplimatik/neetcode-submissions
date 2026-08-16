func search(nums []int, target int) int {
	secondSubArrayStart := 0
	for l, r := 0, len(nums)-1; l <= r; {
		m := (l + r) / 2
		if nums[m] < nums[r] {
			r = m - 1
		} else {
			l = m + 1
		}
		if nums[m] < nums[secondSubArrayStart] {
			secondSubArrayStart = m
		}
	}
	l, r := secondSubArrayStart, len(nums)-1
	if !(nums[l] <= target && target <= nums[r]) {
		l = 0
		r = secondSubArrayStart - 1
	}
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}