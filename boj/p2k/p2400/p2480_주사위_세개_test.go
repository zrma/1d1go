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

func TestSolve2480(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2480")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"3 3 6", "1300"},
		{"6 3 3", "1300"},
		{"2 2 2", "12000"},
		{"6 2 5", "600"},
		{"2 6 5", "600"},
		{"2 5 6", "600"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p2400.Solve2480(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
