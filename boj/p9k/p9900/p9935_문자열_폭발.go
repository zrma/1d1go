package p9900

import (
	"fmt"
)

func Solve9935(reader Reader, writer Writer) {
	var s, bomb string
	_, _ = fmt.Fscan(reader, &s, &bomb)

	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		if len(stack) >= len(bomb) {
			if string(stack[len(stack)-len(bomb):]) == bomb {
				stack = stack[:len(stack)-len(bomb)]
			}
		}
	}

	if len(stack) == 0 {
		//goland:noinspection SpellCheckingInspection
		_, _ = fmt.Fprint(writer, "FRULA")
	} else {
		_, _ = fmt.Fprint(writer, string(stack))
	}
}
