package p1700

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1769(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1769")

	tests := []struct {
		give string
		want string
	}{
		{
			`1234567`,
			`3
NO`,
		},
		{
			`28`,
			`2
NO`,
		},
		{
			`123456`,
			`2
YES`,
		},
		{
			`1`,
			`0
NO`,
		},
		{
			`9`,
			`0
YES`,
		},
		{
			`18`,
			`1
YES`,
		},
		{
			`48`,
			`2
YES`,
		},
		{
			`38`,
			`2
NO`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1769(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
