package p1500

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1541(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1541")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"55-50+40", "-35"},
		{"10+20+30+40", "100"},
		{"00009-00009", "0"},
		{"10-0-0+1", "9"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1541(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
