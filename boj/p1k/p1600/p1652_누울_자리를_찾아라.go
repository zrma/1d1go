package p1600

import (
	"fmt"
	"io"
)

func Solve1652(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscanln(reader, &n)

	board := make([][]rune, n)
	for i := 0; i < n; i++ {
		var s string
		_, _ = fmt.Fscanln(reader, &s)
		board[i] = []rune(s)
	}

	var rowCnt, colCnt int
	for i := 0; i < n; i++ {
		var cnt int
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				cnt++
			} else {
				if cnt >= 2 {
					rowCnt++
				}
				cnt = 0
			}
		}
		if cnt >= 2 {
			rowCnt++
		}
	}
	for j := 0; j < n; j++ {
		var cnt int
		for i := 0; i < n; i++ {
			if board[i][j] == '.' {
				cnt++
			} else {
				if cnt >= 2 {
					colCnt++
				}
				cnt = 0
			}
		}
		if cnt >= 2 {
			colCnt++
		}
	}
	_, _ = fmt.Fprint(writer, rowCnt, colCnt)
}
