package p2000

import (
	"fmt"
	"strings"
)

func Solve2083(reader Reader, writer Writer) {
	for {
		const delim = '\n'
		line, _ := reader.ReadString(delim)
		if line == "# 0 0" || line == "# 0 0\n" {
			break
		}

		var name string
		var age, weight int
		_, _ = fmt.Fscan(strings.NewReader(line), &name, &age, &weight)

		if age > 17 || weight >= 80 {
			_, _ = fmt.Fprintln(writer, name, "Senior")
		} else {
			_, _ = fmt.Fprintln(writer, name, "Junior")
		}
	}
}
