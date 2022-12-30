package p4300_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p4k/p4300"
)

func TestSolve4386(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/4386")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3
1.0 1.0
2.0 2.0
2.0 4.0`,
			"3.41",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p4300.Solve4386(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
