package p1000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1085(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1085")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"6 2 10 3", "1"},
		{"1 1 5 5", "1"},
		{"653 375 1000 1000", "347"},
		{"161 181 762 375", "161"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1085(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
