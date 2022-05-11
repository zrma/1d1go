package p10900

import (
	"bufio"
	"fmt"
	"os"
)

func Solve10952(arr2D [][2]int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() { _ = w.Flush() }()

	for _, v := range arr2D {
		_, _ = fmt.Fprintln(w, v[0]+v[1])
	}
}
