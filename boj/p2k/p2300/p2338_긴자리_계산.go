package p2300

import (
	"fmt"
	"math/big"
)

func Solve2338(reader Reader, writer Writer) {
	var a, b big.Int
	_, _ = fmt.Fscan(reader, &a, &b)

	res := new(big.Int)
	_, _ = fmt.Fprintln(writer, res.Add(&a, &b))
	_, _ = fmt.Fprintln(writer, res.Sub(&a, &b))
	_, _ = fmt.Fprintln(writer, res.Mul(&a, &b))
}
