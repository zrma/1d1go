package p17200

import (
	"fmt"
	"io"
)

func Solve17219(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	passwords := make(map[string]string)
	for i := 0; i < n; i++ {
		var site, password string
		_, _ = fmt.Fscan(reader, &site, &password)
		passwords[site] = password
	}

	for i := 0; i < m; i++ {
		var site string
		_, _ = fmt.Fscan(reader, &site)
		_, _ = fmt.Fprintln(writer, passwords[site])
	}
}
