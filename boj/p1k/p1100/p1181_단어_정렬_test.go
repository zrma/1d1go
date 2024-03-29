package p1100

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1181(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1181")

	tests := []struct {
		give string
		want string
	}{
		{
			`13
but
i
wont
hesitate
no
more
no
more
it
cannot
wait
im
yours`,
			`i
im
it
no
but
more
wait
wont
yours
cannot
hesitate
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1181(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
