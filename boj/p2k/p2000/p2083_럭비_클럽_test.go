package p2000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2083(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2083")

	tests := []struct {
		give string
		want string
	}{
		{
			`Joe 16 34
Bill 18 65
Billy 17 65
Sam 17 85
# 0 0`,
			`Joe Junior
Bill Senior
Billy Junior
Sam Senior
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2083(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
