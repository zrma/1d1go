package p2400_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2400"
)

func TestSolve2482(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2482")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`4
2`,
			"2",
		},
		{
			`20
10`,
			"2",
		},
		{
			`20
11`,
			"0",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p2400.Solve2482(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
