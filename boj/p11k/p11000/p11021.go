package p11000

import (
	"fmt"
	"strconv"
)

func Solve11021(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr2D := make([][2]int, n)
	for i := range arr2D {
		scanner.Scan()
		arr2D[i][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		arr2D[i][1], _ = strconv.Atoi(scanner.Text())
	}

	for i, arr := range arr2D {
		_, _ = fmt.Fprintf(writer, "Case #%d: %d\n", i+1, arr[0]+arr[1])
	}
}
