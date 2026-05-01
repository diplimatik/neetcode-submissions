func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	for i := 1; i < len(res); i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	suf := 1
	for i := len(res) - 1; i >= 0; i-- {
		res[i] *= suf
		suf *= nums[i]
	}
	return res
}