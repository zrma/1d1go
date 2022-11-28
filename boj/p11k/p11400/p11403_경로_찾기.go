package p11400

import (
	"fmt"
)

func Solve11403(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	adj := make([][]int, n)
	for from := 0; from < n; from++ {
		adj[from] = make([]int, n)
		for to := 0; to < n; to++ {
			_, _ = fmt.Fscan(reader, &adj[from][to])
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if adj[i][k] == 1 && adj[k][j] == 1 {
					adj[i][j] = 1
				}
			}
		}
	}

	for from := 0; from < n; from++ {
		for to := 0; to < n; to++ {
			_, _ = fmt.Fprint(writer, adj[from][to])
			if to != n-1 {
				_, _ = fmt.Fprint(writer, " ")
			}
		}
		_, _ = fmt.Fprintln(writer)
	}
}
