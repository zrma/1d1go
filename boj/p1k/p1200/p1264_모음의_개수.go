package p1200

import (
	"fmt"
)

func Solve1264(reader Reader, writer Writer) {
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
