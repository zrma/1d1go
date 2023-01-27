package p1900

import (
	"fmt"
	"io"
)

func Solve1918(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscan(reader, &s)

	// infix to postfix
	stack := make([]byte, 0, len(s))
	res := make([]byte, 0, len(s))
	for _, c := range s {
		switch c {
		case '(':
			stack = append(stack, '(')
		case ')':
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				res = append(res, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		case '+', '-', '*', '/':
			for len(stack) > 0 && priority(stack[len(stack)-1]) >= priority(byte(c)) {
				res = append(res, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, byte(c))
		default:
			res = append(res, byte(c))
		}
	}
	for len(stack) > 0 {
		res = append(res, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	_, _ = fmt.Fprint(writer, string(res))
}

func priority(c byte) int {
	switch c {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	default:
		return 0
	}
}
