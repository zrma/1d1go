package p12900

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve12904(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/12904")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`B
ABBA`,
			`1`,
		},
		{
			`AB
ABB`,
			"0",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve12904(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
