package p1100

import (
	"fmt"
	"strconv"
)

func Solve1193(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	sum := 0
	i := 1
	for ; sum < n; i++ {
		sum += i
	}

	if i%2 == 0 {
		j := sum - n
		_, _ = fmt.Fprintf(writer, "%d/%d", j+1, i-j-1)
	} else {
		j := sum - n
		_, _ = fmt.Fprintf(writer, "%d/%d", i-j-1, j+1)
	}
}
