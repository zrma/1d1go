package p4900

import (
	"bufio"
	"fmt"
	"io"
)

func Solve4949(reader *bufio.Reader, writer io.Writer) {
	for {
		const delim = '\n'
		line, _ := reader.ReadString(delim)
		if line == "." || line == ".\n" {
			break
		}

		if checkBalancedParenthesis(line) {
			_, _ = fmt.Fprintln(writer, "yes")
		} else {
			_, _ = fmt.Fprintln(writer, "no")
		}
	}
}

func checkBalancedParenthesis(s string) bool {
	stack := make([]rune, 0, len(s))
	for _, c := range s {
		if c == '(' || c == '[' {
			stack = append(stack, c)
		} else if c == ')' {
			var ok bool
			stack, ok = pop(stack, '(')
			if !ok {
				return false
			}
		} else if c == ']' {
			var ok bool
			stack, ok = pop(stack, '[')
			if !ok {
				return false
			}
		}
	}
	return len(stack) == 0
}

func pop(stack []rune, comparator rune) ([]rune, bool) {
	if len(stack) == 0 {
		return stack, false
	}
	top := stack[len(stack)-1]
	if top != comparator {
		return stack, false
	}
	return stack[:len(stack)-1], true
}
