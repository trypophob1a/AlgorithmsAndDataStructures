package noddivide

func getNodBySubtraction(a, b int) int {
	if a == b {
		return a
	}

	if a > b {
		return getNodBySubtraction(a-b, b)
	}

	return getNodBySubtraction(b-a, a)
}

func getNodByMod(a, b int) int {
	if a == 0 || b == 0 {
		return b
	}

	if a > b {
		return getNodByMod(a%b, b)
	}

	return getNodByMod(b%a, a)
}
