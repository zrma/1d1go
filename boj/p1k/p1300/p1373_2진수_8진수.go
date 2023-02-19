package p1300

import (
	"fmt"
	"io"
)

func Solve1373(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscanln(reader, &s)

	n := len(s)
	if n%3 == 1 {
		s = "00" + s
	} else if n%3 == 2 {
		s = "0" + s
	}

	n = len(s)
	for i := 0; i < n; i += 3 {
		_, _ = fmt.Fprint(writer, (s[i]-'0')*4+(s[i+1]-'0')*2+(s[i+2]-'0'))
	}
}
