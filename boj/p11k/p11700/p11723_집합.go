package p11700

import (
	"fmt"
	"io"
)

func Solve11723(reader io.Reader, writer io.Writer) {
	var m int
	_, _ = fmt.Fscan(reader, &m)

	s := set{}

	for i := 0; i < m; i++ {
		var command string
		_, _ = fmt.Fscan(reader, &command)

		switch command {
		case "add":
			var x int
			_, _ = fmt.Fscan(reader, &x)
			s.add(x)
		case "remove":
			var x int
			_, _ = fmt.Fscan(reader, &x)
			s.remove(x)
		case "check":
			var x int
			_, _ = fmt.Fscan(reader, &x)
			if s.contains(x) {
				_, _ = fmt.Fprintln(writer, 1)
			} else {
				_, _ = fmt.Fprintln(writer, 0)
			}
		case "toggle":
			var x int
			_, _ = fmt.Fscan(reader, &x)
			if s.contains(x) {
				s.remove(x)
			} else {
				s.add(x)
			}
		case "all":
			s = set{flag: 0b11111111111111111111111111}
		case "empty":
			s = set{flag: 0}

		}
	}
}

type set struct {
	flag int
}

func (s *set) add(x int) {
	s.flag |= 1 << (x - 1)
}

func (s *set) remove(x int) {
	s.flag &= ^(1 << (x - 1))
}

func (s *set) contains(x int) bool {
	return s.flag&(1<<(x-1)) != 0
}
