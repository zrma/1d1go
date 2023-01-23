package p10900

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve10998(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10998")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"1 1", "1"},
		{"1 2", "2"},
		{"1 9", "9"},
		{"9 1", "9"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve10998(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
