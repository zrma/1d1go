package p2900

import (
	"fmt"
	"strings"
)

func Solve2941(reader Reader, writer Writer) {
	const delim = '\n'
	line, _ := reader.ReadString(delim)
	s := strings.Trim(line, "\n")
	s = strings.Trim(s, "\r")

	if len(s) < 2 {
		_, _ = fmt.Fprint(writer, len(s))
		return
	}
	res := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case 'c':
			if i+1 < len(s) && (s[i+1] == '=' || s[i+1] == '-') {
				i++
			}
		case 'd':
			if i+1 < len(s) && s[i+1] == '-' {
				i++
			} else if i+2 < len(s) && (s[i+1] == 'z' && s[i+2] == '=') {
				i += 2
			}
		case 'l':
			if i+1 < len(s) && s[i+1] == 'j' {
				i++
			}
		case 'n':
			if i+1 < len(s) && s[i+1] == 'j' {
				i++
			}
		case 's':
			if i+1 < len(s) && s[i+1] == '=' {
				i++
			}
		case 'z':
			if i+1 < len(s) && s[i+1] == '=' {
				i++
			}
		}
		res++
	}
	_, _ = fmt.Fprint(writer, res)
}
