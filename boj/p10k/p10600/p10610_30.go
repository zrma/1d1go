package p10600

import (
	"fmt"
	"io"
)

func Solve10610(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscan(reader, &s)

	var cnt [10]int
	var sum int64
	for _, c := range s {
		n := int64(c - '0')
		cnt[n]++
		sum += n
	}
	if cnt[0] == 0 || sum%3 != 0 {
		_, _ = fmt.Fprint(writer, -1)
		return
	}
	for i := 9; i >= 0; i-- {
		for j := 0; j < cnt[i]; j++ {
			_, _ = fmt.Fprint(writer, i)
		}
	}
}
