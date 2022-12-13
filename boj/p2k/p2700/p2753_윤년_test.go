package p2700_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2700"
)

func TestSolve2753(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2753")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"2000", "1"},
		{"1999", "0"},
		{"2100", "0"},
		{"2104", "1"},
		{"2200", "0"},
		{"2400", "1"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p2700.Solve2753(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
