package p10800

import (
	"fmt"
	"io"
)

func Solve10828(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	stack := newStack(n)
	for i := 0; i < n; i++ {
		var s string
		_, _ = fmt.Fscan(reader, &s)

		switch s {
		case "push":
			var v int
			_, _ = fmt.Fscan(reader, &v)
			stack.push(v)
		case "pop":
			_, _ = fmt.Fprintln(writer, stack.pop())
		case "size":
			_, _ = fmt.Fprintln(writer, stack.size())
		case "empty":
			if stack.empty() {
				_, _ = fmt.Fprintln(writer, 1)
			} else {
				_, _ = fmt.Fprintln(writer, 0)
			}
		case "top":
			_, _ = fmt.Fprintln(writer, stack.top())
		}
	}
}

type stack struct {
	data []int
}

func newStack(size int) *stack {
	return &stack{data: make([]int, 0, size)}
}

func (s *stack) push(v int) {
	s.data = append(s.data, v)
}

func (s *stack) pop() int {
	if s.empty() {
		return -1
	}

	lastIdx := s.size() - 1
	v := s.data[lastIdx]
	s.data = s.data[:lastIdx]
	return v
}

func (s *stack) size() int {
	return len(s.data)
}

func (s *stack) empty() bool {
	return s.size() == 0
}

func (s *stack) top() int {
	if s.empty() {
		return -1
	}
	lastIdx := s.size() - 1
	return s.data[lastIdx]
}
