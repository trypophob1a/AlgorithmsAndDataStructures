package additionaltasks

// Задача: найти первую пару чисел в отсортированном массиве, сумма которых
// равна заданному числу.

// Входные данные: отсортированном массив целых уникальных чисел arr и целое число x.
// Выходные данные: пара чисел, сумма которых равна x,
// или пустой список, если такой пары не существует.

// Пример:
// findPair([1, 2, 3, 4, 5], 7) => [2, 5]
// findPair([1, 2, 3, 4, 5], 10) => [].

// findPair time: O(n) space: O(1).
func findPair(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}

	l, r := 0, len(nums)-1
	for l < r {
		sum := nums[l] + nums[r]

		if sum == target {
			return []int{nums[l], nums[r]}
		}

		if sum < target {
			l++
		} else {
			r--
		}
	}

	return []int{}
}

// findPairWithBinarySearch time: O(n) space: O(1).
func findPairWithBinarySearch(nums []int, target int) []int {
	res := make([]int, 0)

	if len(nums) < 2 {
		return res
	}

	for _, num := range nums {
		if index := binarySearch(nums, target-num); index != -1 && num != nums[index] {
			return []int{num, nums[index]}
		}
	}

	return res
}

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
