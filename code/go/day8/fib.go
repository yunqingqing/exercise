package fib

func fib(num int) int {
	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	}

	fibNMinusOne := 1
	fibNMinusTwo := 0
	fibN := 0
	for i:=2; i <= num; i++ {
		fibN = fibNMinusOne + fibNMinusTwo
		fibNMinusTwo = fibNMinusOne
		fibNMinusOne = fibN
	}
	return fibN
}