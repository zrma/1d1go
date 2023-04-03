package p16900

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve16916(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/16916")

	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		give string
		want string
	}{
		{
			`baekjoon
aek`,
			"1",
		},
		{
			`baekjoon
bak`,
			"0",
		},
		{
			`baekjoon
joo`,
			"1",
		},
		{
			`baekjoon
oone`,
			"0",
		},
		{
			`baekjoon
online`,
			"0",
		},
		{
			`baekjoon
baekjoon`,
			"1",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve16916(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
