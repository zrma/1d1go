package p9400

import (
	"fmt"
)

func Solve9498(reader Reader, writer Writer) {
	var score int
	_, _ = fmt.Fscan(reader, &score)
	_, _ = fmt.Fprint(writer, scoreToGrade(score))
}

func scoreToGrade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}
