package lv0easy

func romanToInt(s string) int {
	var sum int
	var prev byte
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	isEdgeCase := func(curr byte) bool {
		return prev == 'I' && (curr == 'V' || curr == 'X') ||
			prev == 'X' && (curr == 'L' || curr == 'C') ||
			prev == 'C' && (curr == 'D' || curr == 'M')
	}
	for _, curr := range s {
		sum += m[byte(curr)]
		if isEdgeCase(byte(curr)) {
			sum -= 2 * m[prev]
		}
		prev = byte(curr)
	}
	return sum
}
