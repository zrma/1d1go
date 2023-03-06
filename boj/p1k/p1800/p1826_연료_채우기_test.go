package p1800

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1826(t *testing.T) {
	t.Log("https://github.com/golang/go/issues/1826")

	tests := []struct {
		give string
		want string
	}{
		{
			`4
4 4
5 2
11 5
15 10
25 10`,
			"3",
		},
		{
			`4
4 4
5 2
11 5
15 10
25 5`,
			"4",
		},
		{
			`4
4 4
5 2
11 5
15 10
25 4`,
			"-1",
		},
		{
			`4
4 20
5 2
11 5
15 10
25 5`,
			"1",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1826(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
