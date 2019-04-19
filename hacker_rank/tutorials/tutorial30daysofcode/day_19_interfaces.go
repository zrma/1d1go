package tutorial30daysofcode

import "math"

type advancedArithmetic interface {
	divisorSum(int) int
}

type calculator struct {
}

func (calculator) divisorSum(n int) int {
	var sum int
	s := int(math.Sqrt(float64(n)))

	for i := 1; i <= s; i++ {
		if n%i == 0 {
			d := n / i
			if i == d {
				sum += i
			} else {
				sum += i + (n / i)
			}
		}
	}

	return sum
}

func divisorSum(arithmetic advancedArithmetic, n int) int {
	return arithmetic.divisorSum(n)
}
