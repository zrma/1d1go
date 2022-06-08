package go_basic

func fizzBuzz(n int32) {
	var i int32
	for i = 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			_, _ = funcPrintln("FizzBuzz")
		case i%3 == 0:
			_, _ = funcPrintln("Fizz")
		case i%5 == 0:
			_, _ = funcPrintln("Buzz")
		default:
			_, _ = funcPrintln(i)
		}
	}
}
