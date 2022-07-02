package p11000

import (
	"fmt"

	"1d1go/utils/integer"
)

func Solve11053(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int, n)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	res := integer.LongestIncSubSeqLenSquare(arr)
	_, _ = fmt.Fprint(writer, res)
}
