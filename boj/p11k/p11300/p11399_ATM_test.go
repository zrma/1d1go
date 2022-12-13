package p11300_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11300"
)

func TestSolve11399(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11399")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5
3 1 4 3 2`,
			"32",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p11300.Solve11399(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
