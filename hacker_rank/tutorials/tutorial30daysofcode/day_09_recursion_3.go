package tutorial30daysofcode

var m map[int32]int32

func factorial(n int32) int32 {
	if n <= 1 {
		return 1
	}

	if m == nil {
		m = make(map[int32]int32)
	}

	if v, ok := m[n]; ok {
		return v
	}

	result := n * factorial(n-1)
	m[n] = result
	return result
}
