package p12800

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve12891(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/12891")

	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		give string
		want string
	}{
		{
			`9 8
CCTGGATTG
2 0 1 1`,
			"0",
		},
		{
			`4 2
GATA
1 0 0 1`,
			"2",
		},
		{
			`4 2
ATTA
1 0 0 1`,
			"2",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve12891(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
