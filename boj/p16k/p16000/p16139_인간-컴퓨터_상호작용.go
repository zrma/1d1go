package p16000

import (
	"fmt"
	"io"
)

func Solve16139(reader io.Reader, writer io.Writer) {
	var s string
	var n int
	_, _ = fmt.Fscan(reader, &s, &n)

	const maxLen = 200_000 + 1
	arr := make([][maxLen]int, 26)

	for i := 0; i < len(s); i++ {
		for j := 0; j < 26; j++ {
			arr[j][i+1] = arr[j][i]

			if int(s[i]-'a') == j {
				arr[j][i+1]++
			}
		}
	}

	for i := 0; i < n; i++ {
		var s string
		var begin, end int
		_, _ = fmt.Fscan(reader, &s, &begin, &end)
		char := int(s[0] - 'a')

		res := arr[char][end+1] - arr[char][begin]
		_, _ = fmt.Fprintln(writer, res)
	}
}
