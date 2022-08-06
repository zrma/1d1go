package p10800

import (
	"fmt"
)

func Solve10845(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	q := new(queue)
	for i := 0; i < n; i++ {
		var s string
		_, _ = fmt.Fscan(reader, &s)
		switch s {
		case "push":
			var v int
			_, _ = fmt.Fscan(reader, &v)
			q.push(v)
		case "pop":
			_, _ = fmt.Fprintln(writer, q.pop())
		case "size":
			_, _ = fmt.Fprintln(writer, q.size())
		case "empty":
			if q.empty() {
				_, _ = fmt.Fprintln(writer, 1)
			} else {
				_, _ = fmt.Fprintln(writer, 0)
			}
		case "front":
			_, _ = fmt.Fprintln(writer, q.front())
		case "back":
			_, _ = fmt.Fprintln(writer, q.back())
		}
	}
}

type queue struct {
	data []int
}

func (q *queue) push(v int) {
	q.data = append(q.data, v)
}

func (q *queue) pop() int {
	if q.empty() {
		return -1
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

func (q *queue) size() int {
	return len(q.data)
}

func (q *queue) empty() bool {
	return q.size() == 0
}

func (q *queue) front() int {
	if q.empty() {
		return -1
	}
	return q.data[0]
}

func (q *queue) back() int {
	if q.empty() {
		return -1
	}
	return q.data[len(q.data)-1]
}
