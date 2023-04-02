package p11200

import (
	"fmt"
	"io"
)

func Solve11283(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscanln(reader, &s)

	first, second, third := decomposeKoreanAlphabet([]rune(s)[0])

	_, _ = fmt.Fprint(writer, first*21*28+second*28+third+1)
}

func decomposeKoreanAlphabet(r rune) (int, int, int) {
	s := int(r) - 0xAC00
	return s / (21 * 28), (s / 28) % 21, s % 28
}
