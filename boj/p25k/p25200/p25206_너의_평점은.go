package p25200

import (
	"bufio"
	"fmt"
	"io"
)

func Solve25206(scanner *bufio.Scanner, writer io.Writer) {
	scanner.Split(bufio.ScanLines)

	cnt := 0.0
	sum := 0.0
	for scanner.Scan() {
		s := scanner.Text()

		var lecture string
		var score float64
		var grade string
		_, _ = fmt.Sscanf(s, "%s %f %s", &lecture, &score, &grade)

		if grade == "P" {
			continue
		}

		cnt += score
		switch grade {
		case "A+":
			sum += score * 4.5
		case "A0":
			sum += score * 4.0
		case "B+":
			sum += score * 3.5
		case "B0":
			sum += score * 3.0
		case "C+":
			sum += score * 2.5
		case "C0":
			sum += score * 2.0
		case "D+":
			sum += score * 1.5
		case "D0":
			sum += score * 1.0
		}
	}

	_, _ = fmt.Fprintf(writer, "%.6f", sum/cnt)
}
