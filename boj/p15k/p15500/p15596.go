package p15500

import (
	"fmt"
	"strconv"
)

func Solve15596(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n)
	for i := range arr {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	res := sum(arr)
	_, _ = fmt.Fprint(writer, res)
}

func sum(arr []int) int {
	var res int
	for _, n := range arr {
		res += n
	}
	return res
}
