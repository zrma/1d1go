package p5400

import (
	"fmt"
	"io"
	"strings"
)

func Solve5430(reader io.Reader, writer io.Writer) {
	var t int
	_, _ = fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		processAC(reader, writer)
	}
}

func processAC(reader io.Reader, writer io.Writer) {
	var code string
	_, _ = fmt.Fscan(reader, &code)
	var n int
	_, _ = fmt.Fscan(reader, &n)
	var input string
	_, _ = fmt.Fscan(reader, &input)
	input = input[1 : len(input)-1]

	ss := strings.Split(input, ",")
	arr := make([]string, 0, len(ss))
	for _, s := range ss {
		if s != "" {
			arr = append(arr, s)
		}
	}

	reverse := false
	for _, c := range code {
		switch c {
		case 'R':
			reverse = !reverse
		case 'D':
			if len(arr) == 0 {
				_, _ = fmt.Fprintln(writer, "error")
				return
			}
			if reverse {
				arr = arr[:len(arr)-1]
			} else {
				arr = arr[1:]
			}
		}
	}
	if reverse {
		reverseString(arr)
	}
	_, _ = fmt.Fprintf(writer, "[%s]\n", strings.Join(arr, ","))
}

func reverseString(arr []string) {
	for i, j := 0, len(arr)-1; i < len(arr)/2; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
