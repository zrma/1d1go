package p1200

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1264(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1264")

	tests := []struct {
		give string
		want string
	}{
		{
			`How are you today?
Quite well, thank you, how about yourself?
I live at number twenty four.
#`,
			`7
14
9
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1264(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
