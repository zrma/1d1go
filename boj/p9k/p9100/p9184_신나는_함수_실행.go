package p9100

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve9184(reader Reader, writer Writer) {
	const maxLen = 20 + 1
	cache := make([][][]int, maxLen)
	for i := range cache {
		cache[i] = make([][]int, maxLen)
		for j := range cache[i] {
			cache[i][j] = make([]int, maxLen)
			for k := range cache[i][j] {
				if i == 0 || j == 0 || k == 0 {
					cache[i][j][k] = 1
				} else if i < j && j < k {
					cache[i][j][k] = cache[i][j][k-1] + cache[i][j-1][k-1] - cache[i][j-1][k]
				} else {
					cache[i][j][k] = cache[i-1][j][k] + cache[i-1][j-1][k] + cache[i-1][j][k-1] - cache[i-1][j-1][k-1]
				}
			}
		}
	}

	for {
		const delim = '\n'
		line, _ := reader.ReadString(delim)
		s := strings.Trim(line, "\n")
		s = strings.Trim(s, "\r")

		ss := strings.Split(s, " ")
		a, _ := strconv.Atoi(ss[0])
		b, _ := strconv.Atoi(ss[1])
		c, _ := strconv.Atoi(ss[2])

		if a == -1 && b == -1 && c == -1 {
			break
		}

		res := w(a, b, c, cache)
		_, _ = fmt.Fprintf(writer, "w(%d, %d, %d) = %d\n", a, b, c, res)
	}
}

func w(a, b, c int, cache [][][]int) int {
	if a <= 0 || b <= 0 || c <= 0 {
		return 1
	}
	if a > 20 || b > 20 || c > 20 {
		return cache[20][20][20]
	}
	return cache[a][b][c]
}
