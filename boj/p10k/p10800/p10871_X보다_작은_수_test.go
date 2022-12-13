package p10800_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10800"
)

func TestSolve10871(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10871")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`10 5
1 10 4 9 2 3 8 5 7 6`,
			"1 4 2 3 ",
		},
		{
			`10 4
1 10 4 9 2 3 8 5 7 6`,
			"1 2 3 ",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p10800.Solve10871(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
