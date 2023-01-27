package p1100

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Solve1157(reader *bufio.Reader, writer io.Writer) {
	arr := [26]int{}

	const delim = '\n'
	line, _ := reader.ReadString(delim)
	s := strings.Trim(line, "\n")

	if len(s) == 1 {
		_, _ = fmt.Fprint(writer, strings.ToUpper(s))
		return
	}

	const diff = 'a' - 'A'
	for _, c := range s {
		if c > 'Z' {
			c -= diff
		}
		arr[c-'A']++
	}

	var maxCnt = -1
	var maxChar rune
	for idx, cnt := range arr {
		if cnt > maxCnt {
			maxCnt = cnt
			maxChar = rune(idx + 'A')
		} else if cnt == maxCnt {
			maxChar = '?'
		}
	}
	_, _ = fmt.Fprint(writer, string(maxChar))
}
