package p1700

import (
	"container/heap"
	"fmt"

	"1d1go/utils/ds"
)

func Solve1781(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	toSolve := make(pq1781, 0, n)
	for i := 0; i < n; i++ {
		var d, c int
		_, _ = fmt.Fscan(reader, &d, &c)
		heap.Push(&toSolve, problem{d, c})
	}

	solved := ds.NewMinHeap(n)
	for toSolve.Len() > 0 {
		p := heap.Pop(&toSolve).(problem)
		if p.deadline > solved.Size() {
			solved.Push(p.cupNoodle)
			continue
		}
		if v, _ := solved.Peek(); p.cupNoodle > v {
			solved.Pop()
			solved.Push(p.cupNoodle)
		}
	}

	res := 0
	for solved.Size() > 0 {
		v, _ := solved.Pop()
		res += v
	}
	_, _ = fmt.Fprint(writer, res)
}

type problem struct {
	deadline  int
	cupNoodle int
}

var _ heap.Interface = (*pq1781)(nil)

type pq1781 []problem

func (pq *pq1781) Len() int {
	return len(*pq)
}

func (pq *pq1781) Less(i, j int) bool {
	if (*pq)[i].deadline == (*pq)[j].deadline {
		return (*pq)[i].cupNoodle > (*pq)[j].cupNoodle
	}
	return (*pq)[i].deadline < (*pq)[j].deadline
}

func (pq *pq1781) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *pq1781) Push(x any) {
	*pq = append(*pq, x.(problem))
}

func (pq *pq1781) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
