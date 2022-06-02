package p1100

import (
	"fmt"
	"strings"
)

func Solve1152(reader Reader, writer Writer) {
	const delim = '\n'
	line, _ := reader.ReadString(delim)

	cnt := 0
	for _, s := range strings.Split(line, " ") {
		if s != "" && s != "\n" {
			cnt++
		}
	}
	_, _ = fmt.Fprint(writer, cnt)
}
