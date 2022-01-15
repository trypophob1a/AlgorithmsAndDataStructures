package palindrome

import "unicode"

func isPalindrome(str string) bool {
	if str == "" {
		return true
	}
	var closure func(str string, index int) bool

	runes := []rune(str)
	middle := len(runes) / 2

	closure = func(str string, index int) bool {
		last := (len(runes) - 1) - index

		if middle == index {
			return true
		}
		if unicode.ToLower(runes[last]) != unicode.ToLower(runes[index]) {
			return false
		}

		return closure(str, index+1)
	}
	return closure(str, 0)
}
