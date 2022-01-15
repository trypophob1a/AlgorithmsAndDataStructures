package factorial

func factorial(number int) int {
	if number < 1 {
		return 1
	}

	return number * factorial(number-1)
}

func factorialWithTailRecursion(number, last int) int {
	if number < 1 {
		return last
	}

	return factorialWithTailRecursion(number-1, number*last)
}
