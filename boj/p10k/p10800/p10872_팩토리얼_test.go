package p10800_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10800"
)

func TestSolve10872(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10872")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"10", "3628800"},
		{"0", "1"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p10800.Solve10872(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
