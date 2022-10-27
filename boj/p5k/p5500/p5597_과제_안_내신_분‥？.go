package p5500

import (
	"fmt"
)

func Solve5597(reader Reader, writer Writer) {
	var arr [30]bool
	for i := 0; i < 28; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)
		arr[v-1] = true
	}

	for i, v := range arr {
		if !v {
			_, _ = fmt.Fprintln(writer, i+1)
		}
	}
}
