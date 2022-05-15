package p4300

import (
	"fmt"
	"strconv"
)

func Solve4344(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		countOverAvg(scanner, writer)
	}
}

func countOverAvg(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n)
	sum := 0
	for i := range arr {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
		sum += arr[i]
	}

	avg := float64(sum) / float64(n)
	count := 0
	for _, v := range arr {
		if float64(v) > avg {
			count++
		}
	}

	_, _ = fmt.Fprintf(writer, "%.3f%%\n", float64(count*100)/float64(n))
}
