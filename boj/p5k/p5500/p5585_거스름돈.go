package p5500

import "fmt"

func Solve5585(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	n = 1000 - n

	coins := [6]int{500, 100, 50, 10, 5, 1}

	res := 0
	for _, c := range coins {
		res += n / c
		n %= c
	}
	_, _ = fmt.Fprint(writer, res)
}
