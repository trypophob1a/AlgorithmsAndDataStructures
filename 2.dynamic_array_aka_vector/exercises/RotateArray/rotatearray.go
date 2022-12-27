package rotatearray

// rotate O(n) space O(1).
func rotate(nums []int, k int) {
	k %= len(nums)

	if len(nums) < 2 || k == 0 || k == len(nums) {
		return
	}

	lastIndex := len(nums) - 1

	reverse(nums, 0, lastIndex)
	reverse(nums, 0, k-1)
	reverse(nums, k, lastIndex)
}

// rotate O(n) space O(1).
func reverse(slice []int, begin, end int) {
	for ; begin < end; begin, end = begin+1, end-1 {
		slice[begin], slice[end] = slice[end], slice[begin]
	}
}
