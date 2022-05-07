package card

func reverseString(b []byte) {
	if len(b) == 0 {
		return
	}
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func reverseStringRecur(b []byte) {
	if len(b) < 2 {
		return
	}
	reverseStringFromTo(b, 0, len(b)-1)
}

func reverseStringFromTo(b []byte, i, j int) {
	if i >= j {
		return
	}
	b[i], b[j] = b[j], b[i]
	reverseStringFromTo(b, i+1, j-1)
}
