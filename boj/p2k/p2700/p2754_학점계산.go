package p2700

import (
	"fmt"
	"io"
)

func Solve2754(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscan(reader, &s)

	switch s {
	case "A+":
		_, _ = fmt.Fprintf(writer, "%0.1f", 4.3)
	case "A0":
		_, _ = fmt.Fprintf(writer, "%0.1f", 4.0)
	case "A-":
		_, _ = fmt.Fprintf(writer, "%0.1f", 3.7)
	case "B+":
		_, _ = fmt.Fprintf(writer, "%0.1f", 3.3)
	case "B0":
		_, _ = fmt.Fprintf(writer, "%0.1f", 3.0)
	case "B-":
		_, _ = fmt.Fprintf(writer, "%0.1f", 2.7)
	case "C+":
		_, _ = fmt.Fprintf(writer, "%0.1f", 2.3)
	case "C0":
		_, _ = fmt.Fprintf(writer, "%0.1f", 2.0)
	case "C-":
		_, _ = fmt.Fprintf(writer, "%0.1f", 1.7)
	case "D+":
		_, _ = fmt.Fprintf(writer, "%0.1f", 1.3)
	case "D0":
		_, _ = fmt.Fprintf(writer, "%0.1f", 1.0)
	case "D-":
		_, _ = fmt.Fprintf(writer, "%0.1f", 0.7)
	case "F":
		_, _ = fmt.Fprintf(writer, "%0.1f", 0.0)
	}
}
