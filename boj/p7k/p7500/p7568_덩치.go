package p7500

import (
	"fmt"
	"strconv"
)

func Solve7568(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([][2]int, n)
	for i := range arr {
		scanner.Scan()
		arr[i][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		arr[i][1], _ = strconv.Atoi(scanner.Text())
	}

	for i, v0 := range arr {
		cnt := 1
		for j, v1 := range arr {
			if i == j {
				continue
			}
			if v0[0] < v1[0] && v0[1] < v1[1] {
				cnt++
			}
		}
		_, _ = fmt.Fprintf(writer, "%d ", cnt)
	}
}
