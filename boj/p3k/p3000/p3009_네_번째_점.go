package p3000

import (
	"fmt"
	"strconv"
)

func Solve3009(scanner Scanner, writer Writer) {
	resX := 0
	resY := 0
	for i := 0; i < 3; i++ {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		y, _ := strconv.Atoi(scanner.Text())
		resX ^= x
		resY ^= y
	}
	_, _ = fmt.Fprintf(writer, "%d %d", resX, resY)
}
