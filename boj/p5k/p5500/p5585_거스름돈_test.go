package p5500_test

import (
	"1d1go/boj/p5k/p5500"
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestSolve5585(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/5585")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"380", "4"},
		{"1", "15"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p5500.Solve5585(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
