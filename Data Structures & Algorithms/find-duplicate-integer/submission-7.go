func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[slow]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	head := nums[0]
	for head != nums[slow] {
		head = nums[head]
		slow = nums[slow]
	}
	return head
}