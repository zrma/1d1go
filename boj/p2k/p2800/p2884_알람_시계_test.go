package p2800

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2884(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2884")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"10 10", "9 25"},
		{"0 30", "23 45"},
		{"23 40", "22 55"},
		{"0 45", "0 0"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2884(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
