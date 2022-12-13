package p13300_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p13k/p13300"
)

func TestSolve13305(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/13305")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`4
2 3 1
5 2 4 1`,
			"18",
		},
		{
			`4
3 3 4
1 1 1 1`,
			"10",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p13300.Solve13305(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
