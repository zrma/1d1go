package p1300

import (
	"fmt"
	"io"
)

func Solve1356(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscanln(reader, &s)

	if len(s) == 1 {
		_, _ = fmt.Fprint(writer, "NO")
		return
	}

	leftArr := make([]int64, len(s)+1)
	rightArr := make([]int64, len(s)+1)

	leftArr[0] = 1
	rightArr[len(s)] = 1

	for i := 1; i < len(s); i++ {
		leftArr[i] = leftArr[i-1] * int64(s[i-1]-'0')
		rightArr[len(s)-i] = rightArr[len(s)-i+1] * int64(s[len(s)-i]-'0')
	}

	for i := 1; i < len(s); i++ {
		if leftArr[i] == rightArr[i] {
			_, _ = fmt.Fprint(writer, "YES")
			return
		}
	}

	_, _ = fmt.Fprint(writer, "NO")
}
