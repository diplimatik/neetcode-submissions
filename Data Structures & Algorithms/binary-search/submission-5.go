func search(nums []int, target int) int {
	for l, r := 0, len(nums)-1; ; {
		if l == r {
			if nums[l] == target {
				return l
			}
			break
		}
		if nums[(r+l)/2] == target {
			return (r + l) / 2
		} else if nums[(r+l)/2] < target {
			l = (r+l)/2 + 1
		} else {
			r = (r + l) / 2
		}
	}
	return -1
}