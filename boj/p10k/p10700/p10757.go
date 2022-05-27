package p10700

import (
	"fmt"
)

func Solve10757(scanner Scanner, writer Writer) {
	scanner.Scan()
	a := scanner.Text()
	scanner.Scan()
	b := scanner.Text()

	const maxLength = 10000 + 1
	res := [maxLength]int{}
	for i := 0; i < maxLength; i++ {
		if i < len(a) {
			res[i] += int(a[len(a)-i-1] - '0')
		}
		if i < len(b) {
			res[i] += int(b[len(b)-i-1] - '0')
		}
		if res[i] >= 10 {
			res[i+1] += res[i] / 10
			res[i] %= 10
		}
	}

	var found bool
	for i := len(res) - 1; i >= 0; i-- {
		if !found && res[i] == 0 {
			continue
		}
		found = true
		_, _ = fmt.Fprint(writer, res[i])
	}
}
