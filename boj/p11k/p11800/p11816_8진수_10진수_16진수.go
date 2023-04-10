package p11800

import (
	"fmt"
	"io"
	"strings"
)

func Solve11816(reader io.Reader, writer io.Writer) {
	var n string
	_, _ = fmt.Fscanln(reader, &n)

	switch {
	case strings.HasPrefix(n, "0x"):
		_, _ = fmt.Fprint(writer, hexToDex(n[2:]))
	case strings.HasPrefix(n, "0"):
		_, _ = fmt.Fprint(writer, octToDex(n[1:]))
	default:
		_, _ = fmt.Fprint(writer, n)
	}
}

func hexToDex(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		var tmp int
		switch {
		case '0' <= s[i] && s[i] <= '9':
			tmp = int(s[i] - '0')
		case 'a' <= s[i] && s[i] <= 'f':
			tmp = int(s[i]-'a') + 10
		}
		res = res*16 + tmp
	}
	return res
}

func octToDex(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		res = res*8 + int(s[i]-'0')
	}
	return res
}
