package day14

func Pow(base float64, exponent int) float64 {
	if base == 0 {
		return 0
	}
	powFlag := false
	if exponent < 0 {
		powFlag = true
		exponent = -exponent
	}
	result := calPow(base, exponent)

	if powFlag {
		result = 1 / result
	}

	return result
}

func calPow(base float64, exponent int) float64 {
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}
	result := calPow(base, exponent>>1)
	result *= result

	if exponent&1 == 1 {
		result *= base
	}

	return result
}
