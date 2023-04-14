package p1700

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1755(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1755")

	tests := []struct {
		give string
		want string
	}{
		{
			`8 28`,
			`8 9 18 15 14 19 11 17 16 13
12 10 28 25 24 21 27 26 23 22
20`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1755(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
