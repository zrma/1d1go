package p9300

import (
	"fmt"
	"strconv"

	"1d1go/programmers/practice"
)

func Solve9375(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())

		arr := make([][]string, k)
		for j := range arr {
			arr[j] = make([]string, 2)

			scanner.Scan()
			arr[j][0] = scanner.Text()
			scanner.Scan()
			arr[j][1] = scanner.Text()
		}
		res := practice.Disguise(arr)
		_, _ = fmt.Fprintln(writer, res)
	}
}
