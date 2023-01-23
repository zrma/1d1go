package p3000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve3003(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/3003")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"0 1 2 2 2 7", "1 0 0 0 0 1"},
		{"2 1 2 1 2 1", "-1 0 0 1 0 7"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve3003(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
