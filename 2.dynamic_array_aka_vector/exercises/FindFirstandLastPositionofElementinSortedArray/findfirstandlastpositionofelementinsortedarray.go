package findfirstandlastpositionofelementinsortedarray

func searchRange(nums []int, target int) []int {
	return []int{searchFirst(nums, target), searchLast(nums, target)}
}

func searchLast(slice []int, needle int) int {
	last := -1
	l, r := 0, len(slice)-1

	for l <= r {
		m := (l + r) >> 1

		if slice[m] == needle {
			if m < r && slice[m] != slice[m+1] {
				return m
			}
			last = m
		}

		if slice[m] <= needle {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return last
}

func searchFirst(slice []int, needle int) int {
	first := -1
	l, r := 0, len(slice)-1

	for l <= r {
		m := (l + r) >> 1

		if slice[m] == needle {
			if m > 0 && slice[m] != slice[m-1] {
				return m
			}
			first = m
		}

		if slice[m] >= needle {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return first
}
