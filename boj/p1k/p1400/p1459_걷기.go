package p1400

import (
	"fmt"
	"io"
)

func Solve1459(reader io.Reader, writer io.Writer) {
	var x, y, w, s int
	_, _ = fmt.Fscan(reader, &x, &y, &w, &s)

	// 수직, 수평으로만 이동
	if w*2 < s {
		_, _ = fmt.Fprint(writer, (x+y)*w)
		return
	}

	if y < x {
		x, y = y, x
	}

	// 최대한 수직/수평 이동 + 남는 거리 대각선
	if w < s {
		_, _ = fmt.Fprint(writer, x*s+(y-x)*w)
		return
	}

	if (x+y)%2 == 0 {
		// 대각선으로만 이동
		_, _ = fmt.Fprint(writer, y*s)
	} else {
		// 대각선으로만 이동하면 수직/수평 1개 남음
		_, _ = fmt.Fprint(writer, (y-1)*s+w)
	}
}
