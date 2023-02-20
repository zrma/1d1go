package p7500

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve7567(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/7567")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			"((((",
			"25",
		},
		{
			"()()()))(",
			"80",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve7567(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
