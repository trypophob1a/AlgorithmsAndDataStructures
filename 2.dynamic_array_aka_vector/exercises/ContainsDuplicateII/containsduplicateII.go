package containsduplicateii

// naiveContainsNearbyDuplicate O(n^2).
func naiveContainsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := 1 + i; j < len(nums); j++ {
			if nums[i] == nums[j] && j-i <= k {
				return true
			}
		}
	}
	return false
}

// containsNearbyDuplicate O(n) space O(n).
func containsNearbyDuplicate(nums []int, k int) bool {
	memo := make(map[int]int, len(nums))
	for i, num := range nums {
		if val, ok := memo[num]; ok && i-val <= k {
			return true
		}
		memo[num] = i
	}
	return false
}
