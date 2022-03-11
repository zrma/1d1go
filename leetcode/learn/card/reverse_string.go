package card

func reverseString(s []byte) {
	if len(s) == 0 {
		return
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseStringRecur(s []byte) {
	if len(s) < 2 {
		return
	}
	reverseStringFromTo(s, 0, len(s)-1)
}

func reverseStringFromTo(s []byte, i, j int) {
	if i >= j {
		return
	}
	s[i], s[j] = s[j], s[i]
	reverseStringFromTo(s, i+1, j-1)
}
