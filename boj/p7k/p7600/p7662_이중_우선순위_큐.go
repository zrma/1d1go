package p7600

import (
	"fmt"
	"io"

	"1d1go/utils/ds"
)

func Solve7662(reader io.Reader, writer io.Writer) {
	var t int
	_, _ = fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		solve7662(reader, writer)
	}
}

func solve7662(reader io.Reader, writer io.Writer) {
	var k int
	_, _ = fmt.Fscan(reader, &k)

	maxH := ds.NewMaxHeap(k)
	minH := ds.NewMinHeap(k)

	popCountsInMax := make(map[int]int)
	popCountsInMin := make(map[int]int)

	for i := 0; i < k; i++ {
		var op string
		var v int
		_, _ = fmt.Fscan(reader, &op, &v)

		if op == "I" {
			maxH.Push(v)
			minH.Push(v)
			continue
		}

		if v == 1 {
			for {
				pop, ok := maxH.Pop()
				if !ok {
					break
				}

				if cnt, ok := popCountsInMin[pop]; ok {
					if cnt == 1 {
						delete(popCountsInMin, pop)
					} else {
						popCountsInMin[pop]--
					}
					continue
				}

				popCountsInMax[pop]++
				break
			}
		} else {
			for {
				pop, ok := minH.Pop()
				if !ok {
					break
				}

				if cnt, ok := popCountsInMax[pop]; ok {
					if cnt == 1 {
						delete(popCountsInMax, pop)
					} else {
						popCountsInMax[pop]--
					}
					continue
				}

				popCountsInMin[pop]++
				break
			}
		}
	}

	trimResult(popCountsInMax, minH)
	trimResult(popCountsInMin, maxH)

	if maxH.Size() == 0 || minH.Size() == 0 {
		_, _ = fmt.Fprintln(writer, "EMPTY")
		return
	}

	max, _ := maxH.Peek()
	min, _ := minH.Peek()
	_, _ = fmt.Fprintf(writer, "%d %d\n", max, min)
}

type priorityQueue interface {
	Peek() (int, bool)
	Pop() (int, bool)
}

func trimResult(popCounts map[int]int, pq priorityQueue) {
	for len(popCounts) > 0 {
		peek, _ := pq.Peek()

		cnt, ok := popCounts[peek]
		if !ok {
			break
		}
		if cnt == 1 {
			delete(popCounts, peek)
		} else {
			popCounts[peek]--
		}
		pq.Pop()
	}
}
