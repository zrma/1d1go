package p1200

import (
	"fmt"
	"io"
	"math/big"
)

func Solve1271(reader io.Reader, writer io.Writer) {
	var n, m big.Int
	_, _ = fmt.Fscan(reader, &n, &m)

	res := new(big.Int)
	_, _ = fmt.Fprintln(writer, res.Div(&n, &m))
	_, _ = fmt.Fprintln(writer, res.Mod(&n, &m))
}
