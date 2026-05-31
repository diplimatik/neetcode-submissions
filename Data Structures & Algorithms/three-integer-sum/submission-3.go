func threeSum(nums []int) [][]int {
	numsMap := make(map[int]int)
	ansMap := make(map[[3]int]struct{})
	for i, n := range nums {
		numsMap[n] = i
	}

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums); j++ {
			numToFind := -(nums[i] + nums[j])
			if matchIdx, ok := numsMap[numToFind]; ok && matchIdx != i && matchIdx != j {
				a, b, c := nums[i], nums[j], numToFind
				a, b = min(a, b), max(a, b)
				b, c = min(b, c), max(b, c)
				a, b = min(a, b), max(a, b)
				ansMap[[3]int{a, b, c}] = struct{}{}
			}
		}
	}

	ans := make([][]int, 0)
	for triplet, _ := range ansMap {
		ans = append(ans, triplet[:])
	}

	return ans
}