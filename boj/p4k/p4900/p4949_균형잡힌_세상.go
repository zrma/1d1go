package p4900

import (
	"fmt"
)

func Solve4949(reader Reader, writer Writer) {
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
	stack0 := make([]rune, 0, len(s))
	stack1 := make([]rune, 0, len(s))
	for _, c := range s {
		if c == '(' || c == '[' {
			stack0 = append(stack0, c)
		} else if c == ')' {
			if len(stack0) == 0 {
				return false
			}
			pop := stack0[len(stack0)-1]
			if pop != '(' {
				return false
			}
			stack0 = stack0[:len(stack0)-1]
		} else if c == ']' {
			if len(stack0) == 0 {
				return false
			}
			pop := stack0[len(stack0)-1]
			if pop != '[' {
				return false
			}
			stack0 = stack0[:len(stack0)-1]
		}
	}
	return len(stack0) == 0 && len(stack1) == 0
}
