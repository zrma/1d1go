package dp

import (
	"strconv"
)

func substrings(n string) int32 {
	const modulo = 1000000007

	length := len(n)
	prev, err := strconv.Atoi(string(n[0]))
	if err != nil {
		return 0
	}

	sum := prev
	var i int
	for i = 1; i < length; i++ {
		sub, err := strconv.Atoi(string(n[i]))
		if err != nil {
			sub = 0
		}
		sum = (sum*10 + sub*(i+1)) % modulo
		prev = (sum + prev) % modulo
	}

	return int32(prev)
}
