package p1000_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1000"
)

func TestSolve1024(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1024")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"18 2", "5 6 7"},
		{"18 4", "3 4 5 6"},
		{"18 5", "-1"},
		{"45 10", "0 1 2 3 4 5 6 7 8 9"},
		{"1000000000 2", "199999998 199999999 200000000 200000001 200000002"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p1000.Solve1024(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
