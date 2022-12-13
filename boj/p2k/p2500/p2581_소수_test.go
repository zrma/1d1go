package p2500_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2500"
)

func TestSolve2581(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2581")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`60
100`,
			`620
61`,
		},
		{
			`64
65`,
			"-1",
		},
		{
			`1
2`,
			`2
2`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p2500.Solve2581(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
