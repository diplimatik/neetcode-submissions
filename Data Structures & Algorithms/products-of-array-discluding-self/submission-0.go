func productExceptSelf(nums []int) []int {
	pref, suf := make([]int, len(nums)), make([]int, len(nums))
	pref[0] = 1
	suf[len(suf)-1] = 1
	for i := 1; i < len(pref); i++ {
		pref[i] = nums[i-1] * pref[i-1]
	}

	for i := len(suf) - 2; i >= 0; i-- {
		suf[i] = nums[i+1] * suf[i+1]
	}
	res := make([]int, len(nums))
	for i := 0; i < len(res); i++ {
		res[i] = suf[i] * pref[i]
	}
	
	return res
}