func twoSum(numbers []int, target int) []int {
	j := 1
	goRight := true
	for i := 0; i < len(numbers)-1; i++ {
		for j < len(numbers) && j > i {
			if goRight {
				if numbers[i]+numbers[j] >= target {
					if j != i+1 {
						goRight = false
					}
					break
				}
				j++
			} else {
				if numbers[i]+numbers[j] <= target {
					if j != len(numbers)-1 {
						goRight = true
					}
					break
				}
				j--
			}
		}
		if goRight && j == len(numbers) {
			j--
			goRight = false
		} else if !goRight && j == i {
			j++
			goRight = true
		}

		if numbers[i]+numbers[j] == target {
			return []int{min(i, j) + 1, max(i, j) + 1}
		}
	}
	return []int{}
}
