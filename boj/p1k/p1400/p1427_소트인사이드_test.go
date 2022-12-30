package p1400_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1400"
)

func TestSolve1427(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1427")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"2143", "4321"},
		{"999998999", "999999998"},
		{"61423", "64321"},
		{"500613009", "965310000"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p1400.Solve1427(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
