package p1200

import (
	"bufio"
	"fmt"
	"io"
)

func Solve1264(reader *bufio.Reader, writer io.Writer) {
	for {
		const delim = '\n'
		line, _ := reader.ReadString(delim)
		if line == "#" || line == "#\n" {
			break
		}

		var count int
		for _, c := range line {
			switch c {
			case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
				count++
			}
		}
		_, _ = fmt.Fprintln(writer, count)
	}
}
