func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	for i := 1; i < len(res); i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	suf := make([]int, len(nums))
	suf[len(suf)-1] = 1
	for i := len(suf) - 2; i >= 0; i-- {
		suf[i] = suf[i+1] * nums[i+1]
	}

	for n := range res {
		res[n] = suf[n] * res[n]
	}
	return res
}