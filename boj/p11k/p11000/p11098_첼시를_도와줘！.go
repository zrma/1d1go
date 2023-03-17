package p11000

import (
	"fmt"
	"io"
)

func Solve11098(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		var p int
		_, _ = fmt.Fscanln(reader, &p)

		var maxName string
		var maxPrice int
		for j := 0; j < p; j++ {
			var price int
			var name string
			_, _ = fmt.Fscanln(reader, &price, &name)

			if price > maxPrice {
				maxPrice = price
				maxName = name
			}
		}

		_, _ = fmt.Fprintln(writer, maxName)
	}
}
