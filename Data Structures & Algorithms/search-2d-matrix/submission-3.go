func searchMatrix(matrix [][]int, target int) bool {
	arr := make([]int, 0)
	for l, r := 0, len(matrix)-1; l <= r; {
		mid := (l + r) / 2
		if matrix[mid][0] <= target && matrix[mid][len(matrix[0])-1] >= target {
			arr = matrix[mid]
			break
		} else if matrix[mid][0] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	for l, r := 0, len(arr)-1; l <= r; {
		if arr[(l+r)/2] == target {
			return true
		} else if arr[(l+r)/2] > target {
			r = (l+r)/2 - 1
		} else {
			l = (l+r)/2 + 1
		}
	}
	return false
}