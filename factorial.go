package factorial

func CalculateFactorial(num int) int {
	num++
	if num == 1 || num == 0 {
		
		return num
	}

	return num * CalculateFactorial(num-1)
}

type Product struct {
	ID       string
	Name     string
	Quantity int
}
