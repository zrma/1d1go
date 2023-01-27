package p9000

import (
	"fmt"
	"io"
)

func Solve9019(reader io.Reader, writer io.Writer) {
	var t int
	_, _ = fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		solve9019(reader, writer)
	}
}

func solve9019(reader io.Reader, writer io.Writer) {
	var a, b int
	_, _ = fmt.Fscan(reader, &a, &b)

	const max = 10000
	visited := make([]bool, max)
	commands := make([]string, max)

	queue := make([]int, 0, max)
	queue = append(queue, a)
	visited[a] = true

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if p == b {
			break
		}

		d := (p * 2) % max
		if !visited[d] {
			visited[d] = true
			commands[d] = commands[p] + "D"
			queue = append(queue, d)
		}

		s := (p - 1 + max) % max
		if !visited[s] {
			visited[s] = true
			commands[s] = commands[p] + "S"
			queue = append(queue, s)
		}

		l := (p % 1000) * 10
		l += p / 1000
		if !visited[l] {
			visited[l] = true
			commands[l] = commands[p] + "L"
			queue = append(queue, l)
		}

		r := (p % 10) * 1000
		r += p / 10
		if !visited[r] {
			visited[r] = true
			commands[r] = commands[p] + "R"
			queue = append(queue, r)
		}
	}
	_, _ = fmt.Fprintln(writer, commands[b])
}
