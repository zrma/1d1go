package p13700

import (
	"fmt"
	"io"
	"math/big"
)

func Solve13706(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscanln(reader, &s)

	n := &big.Int{}
	n, _ = n.SetString(s, 10)

	sqrt := new(big.Int).Sqrt(n)
	_, _ = fmt.Fprint(writer, sqrt)
}
