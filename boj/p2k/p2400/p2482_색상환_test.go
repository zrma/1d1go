package p2400

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2482(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2482")

	tests := []struct {
		give string
		want string
	}{
		{
			`4
2`,
			"2",
		},
		{
			`20
10`,
			"2",
		},
		{
			`20
11`,
			"0",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2482(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
