package p4400

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Solve4470(reader *bufio.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	for i := 1; i <= n; i++ {
		const delim = '\n'
		line, _ := reader.ReadString(delim)
		s := strings.Trim(line, "\n")
		s = strings.Trim(s, "\r")
		if s == "" {
			i--
			continue
		}

		_, _ = fmt.Fprintf(writer, "%d. %s\n", i, s)
	}
}
