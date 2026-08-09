func search(nums []int, target int) int {
	for l, r := 0, len(nums)-1; l <= r; {
		if nums[(r+l)/2] == target {
			return (r + l) / 2
		} else if nums[(r+l)/2] < target {
			l = (r+l)/2 + 1
		} else {
			r = (r+l)/2 - 1
		}
	}
	return -1
}