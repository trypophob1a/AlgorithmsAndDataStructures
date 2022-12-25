package sortarraybyparity

func sortArrayByParity(nums []int) []int {
	evenIndex := 0
	for i := 0; i < len(nums); i++ {
		if (1 & nums[i]) == 0 {
			nums[evenIndex], nums[i] = nums[i], nums[evenIndex]
			evenIndex++
		}
	}

	return nums
}
