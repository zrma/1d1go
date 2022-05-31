package p18800

import (
	"fmt"
	"sort"
	"strconv"
)

func Solve18870(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	exist := make(map[int]bool)
	set := make([]int, 0, n)
	arr := make([]int, n)
	for i := range arr {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		arr[i] = v
		if _, ok := exist[v]; !ok {
			set = append(set, v)
			exist[v] = true
		}
	}
	sort.Ints(set)

	for _, v := range arr {
		_, _ = fmt.Fprintf(writer, "%d ", sort.SearchInts(set, v))
	}
}
