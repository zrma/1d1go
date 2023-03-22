package p5200

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve5218(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/5218")

	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		give string
		want string
	}{
		{
			`5
AAAA ABCD
ABCD AAAA
DARK LOKI
STRONG THANOS
DEADLY ULTIMO`,
			`Distances: 0 1 2 3
Distances: 0 25 24 23
Distances: 8 14 19 24
Distances: 1 14 9 25 1 12
Distances: 17 7 19 5 1 16
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve5218(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
