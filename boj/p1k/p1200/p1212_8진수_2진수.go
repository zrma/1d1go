package p1200

import (
	"fmt"
	"io"
	"strings"
)

func Solve1212(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscan(reader, &s)

	for i, c := range s {
		if i == 0 {
			v := strings.TrimLeft(toBinary(byte(c)), "0")
			if v == "" {
				_, _ = fmt.Fprint(writer, 0)
			} else {
				_, _ = fmt.Fprint(writer, v)
			}
		} else {
			_, _ = fmt.Fprint(writer, toBinary(byte(c)))
		}
	}
}

func toBinary(c byte) string {
	switch c {
	case '0':
		return "000"
	case '1':
		return "001"
	case '2':
		return "010"
	case '3':
		return "011"
	case '4':
		return "100"
	case '5':
		return "101"
	case '6':
		return "110"
	case '7':
		return "111"
	}
	return ""
}
