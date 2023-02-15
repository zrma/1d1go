package p17400

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Solve17413(scanner *bufio.Scanner, writer io.Writer) {
	scanner.Split(bufio.ScanLines)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	input := scanner.Text()

	res := strings.Builder{}

	stack := make([]byte, 0, len(input))
	flushStack := func() {
		for len(stack) > 0 {
			res.WriteByte(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}

	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '<':
			flushStack()
			for i < len(input) && input[i] != '>' {
				res.WriteByte(input[i])
				i++
			}
			res.WriteByte(input[i])
		case ' ':
			flushStack()
			res.WriteByte(' ')
		default:
			stack = append(stack, input[i])
		}
	}
	flushStack()

	_, _ = fmt.Fprint(writer, res.String())
}
