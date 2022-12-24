package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	//	 v		 	        j = 1
	// 0,0,1,1,2
	// ^      		        i = 0
	// arr 0,0,1,1,2
	// c = 1

	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[count] = nums[i]
			count++
		}
	}

	return count
}
