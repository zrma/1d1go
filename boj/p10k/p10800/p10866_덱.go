package p10800

import (
	"fmt"
	"io"
)

func Solve10866(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	deq := newQueue(n)

	for i := 0; i < n; i++ {
		var s string
		_, _ = fmt.Fscan(reader, &s)
		switch s {
		case "push_front":
			var v int
			_, _ = fmt.Fscan(reader, &v)
			deq.pushFront(v)
		case "push_back":
			var v int
			_, _ = fmt.Fscan(reader, &v)
			deq.pushBack(v)
		case "pop_front":
			_, _ = fmt.Fprintln(writer, deq.popFront())
		case "pop_back":
			_, _ = fmt.Fprintln(writer, deq.popBack())
		case "size":
			_, _ = fmt.Fprintln(writer, deq.size())
		case "empty":
			if deq.empty() {
				_, _ = fmt.Fprintln(writer, 1)
			} else {
				_, _ = fmt.Fprintln(writer, 0)
			}
		case "front":
			_, _ = fmt.Fprintln(writer, deq.front())
		case "back":
			_, _ = fmt.Fprintln(writer, deq.back())
		}
	}
}

func newQueue(n int) *deque {
	return &deque{
		data: make([]int, 0, n),
	}
}

type deque struct {
	data []int
}

func (d *deque) pushFront(v int) {
	d.data = append([]int{v}, d.data...)
}

func (d *deque) pushBack(v int) {
	d.data = append(d.data, v)
}

func (d *deque) popFront() int {
	if d.empty() {
		return -1
	}
	v := d.data[0]
	d.data = d.data[1:]
	return v
}

func (d *deque) popBack() int {
	if d.empty() {
		return -1
	}
	v := d.data[len(d.data)-1]
	d.data = d.data[:len(d.data)-1]
	return v
}

func (d *deque) size() int {
	return len(d.data)
}

func (d *deque) empty() bool {
	return d.size() == 0
}

func (d *deque) front() int {
	if d.empty() {
		return -1
	}
	return d.data[0]
}

func (d *deque) back() int {
	if d.empty() {
		return -1
	}
	return d.data[len(d.data)-1]
}
