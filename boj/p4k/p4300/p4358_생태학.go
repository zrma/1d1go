package p4300

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

func Solve4358(scanner *bufio.Scanner, writer io.Writer) {
	scanner.Split(bufio.ScanLines)

	tree := make(map[string]int)
	for scanner.Scan() {
		tree[scanner.Text()]++
	}

	tot := 0
	for _, v := range tree {
		tot += v
	}

	var keys []string
	for k := range tree {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		_, _ = fmt.Fprintf(writer, "%s %.4f\n", k, float64(tree[k])/float64(tot)*100)
	}
}
