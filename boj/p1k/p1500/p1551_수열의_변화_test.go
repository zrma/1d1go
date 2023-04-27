package p1500

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1551(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1551")

	tests := []struct {
		give string
		want string
	}{
		{
			`5 1
5,6,3,9,-1`,
			`1,-3,6,-10`,
		},
		{
			`5 2
5,6,3,9,-1`,
			`-4,9,-16`,
		},
		{
			`5 4
5,6,3,9,-1`,
			`-38`,
		},
		{
			`8 3
4,4,4,4,4,4,4,4`,
			`0,0,0,0,0`,
		},
		{
			`2 0
-100,100`,
			`-100,100`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1551(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
