package p7800_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p7k/p7800"
)

func TestSolve7869(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/7869")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			"20.0 30.0 15.0 40.0 30.0 30.0",
			"608.366",
		},
		{
			"0.0 0.0 100.0 200.0 0.0 100.0",
			"0.000",
		},
		{
			"0.0 0.0 100.0 0.0 0.0 50.0",
			"7853.982",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p7800.Solve7869(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
