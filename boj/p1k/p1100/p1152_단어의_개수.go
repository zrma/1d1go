package p1100

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Solve1152(reader *bufio.Reader, writer io.Writer) {
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
