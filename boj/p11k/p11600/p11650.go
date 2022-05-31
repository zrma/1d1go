package p11600

import (
	"fmt"
	"sort"
	"strconv"
)

func Solve11650(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([][2]int, n)
	for i := range arr {
		scanner.Scan()
		arr[i][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		arr[i][1], _ = strconv.Atoi(scanner.Text())
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] < arr[j][0] {
			return true
		}
		if arr[i][0] > arr[j][0] {
			return false
		}
		return arr[i][1] < arr[j][1]
	})

	for _, v := range arr {
		_, _ = fmt.Fprintln(writer, v[0], v[1])
	}
}
