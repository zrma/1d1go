package tutorial30daysofcode

import "math"

func isPrime(num int32) bool {
	if num == 1 {
		return false
	}

	if num == 2 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	sqNum := math.Sqrt(float64(num))

	var i int32 = 3
	for ; i < int32(sqNum)+1; i += 2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}
