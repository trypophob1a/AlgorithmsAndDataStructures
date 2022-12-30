package threesum

import "sort"

// threeSum O(n^2) space O(n).
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			switch {
			case sum == 0:
				res = append(res, []int{nums[i], nums[l], nums[r]})

				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}

				for l < r && nums[r] == nums[r+1] {
					r--
				}
			case sum < 0:
				l++
			default:
				r--
			}
		}
	}

	return res
}
