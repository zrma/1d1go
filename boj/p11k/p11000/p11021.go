package p11000

import (
	"bufio"
	"fmt"
	"os"
)

func Solve11021(arr2D [][2]int) {
	w := bufio.NewWriter(os.Stdout)

	for i, arr := range arr2D {
		_, _ = fmt.Fprintf(w, "Case #%d: %d\n", i+1, arr[0]+arr[1])
	}
	_ = w.Flush()
}
