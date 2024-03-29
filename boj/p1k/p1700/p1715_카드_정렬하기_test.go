package p1700

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1715(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1715")

	tests := []struct {
		give string
		want string
	}{
		{
			`3
10
20
40`,
			"100",
		},
		{
			`10
123
456
234
998
12
7
876
887
455
234`,
			"12108",
		},
		{
			`1
100`,
			"0",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1715(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
