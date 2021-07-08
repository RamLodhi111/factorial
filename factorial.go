package factorial

func CalculateFactorial(num int) int {

	if num == 1 || num == 0 {
		return num
	}

	return num * CalculateFactorial(num-1)
}
