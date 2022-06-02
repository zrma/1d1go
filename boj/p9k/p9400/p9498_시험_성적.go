package p9400

import (
	"fmt"
	"strconv"
)

func Solve9498(scanner Scanner, writer Writer) {
	scanner.Scan()
	score, _ := strconv.Atoi(scanner.Text())
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
