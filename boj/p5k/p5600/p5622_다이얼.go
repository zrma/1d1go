package p5600

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Solve5622(reader *bufio.Reader, writer io.Writer) {
	const delim = '\n'
	line, _ := reader.ReadString(delim)
	s := strings.Trim(line, "\n")
	s = strings.Trim(s, "\r")

	arr := [26]int{
		3, 3, 3,
		4, 4, 4,
		5, 5, 5,
		6, 6, 6,
		7, 7, 7,
		8, 8, 8, 8,
		9, 9, 9,
		10, 10, 10, 10,
	}

	var res int32
	for _, c := range s {
		res += int32(arr[c-'A'])
	}
	_, _ = fmt.Fprint(writer, res)
}
