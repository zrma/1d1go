package p11000

import (
	"fmt"
	"strconv"

	"1d1go/utils/integer"
)

func Solve11053(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n)
	for i := range arr {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	res := integer.LongestIncSubSeqLen(arr)
	_, _ = fmt.Fprint(writer, res)
}
