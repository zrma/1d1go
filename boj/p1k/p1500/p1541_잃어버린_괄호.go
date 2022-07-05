package p1500

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve1541(reader Reader, writer Writer) {
	var input string
	_, _ = fmt.Fscan(reader, &input)

	ss := strings.Split(input, "-")

	res := 0
	for i, s0 := range ss {
		sum := 0
		for _, s1 := range strings.Split(s0, "+") {
			v, _ := strconv.Atoi(s1)
			sum += v
		}
		if i == 0 {
			res = sum
		} else {
			res -= sum
		}
	}
	_, _ = fmt.Fprint(writer, res)
}
