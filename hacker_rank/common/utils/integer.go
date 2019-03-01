package utils

func Min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

func Max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

func Pow(n, p int) int {
	if n == 0 {
		return 0
	}

	result := 1
	for ; p > 0; p-- {
		result *= n
	}

	return result
}
