package fibonaccinumbers

func naiveAlgorithm(number uint) uint {
	if number == 0 {
		return 0
	}

	if number < 3 {
		return 1
	}

	return naiveAlgorithm(number-1) + naiveAlgorithm(number-2)
}
