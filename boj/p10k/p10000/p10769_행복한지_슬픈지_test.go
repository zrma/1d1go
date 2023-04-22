package p10000

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve10769(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10769")

	tests := []struct {
		give string
		want string
	}{
		{"How are you :-) doing :-( today :-)?", "happy"},
		{":)", "none"},
		{"This:-(is str:-(:-(ange te:-)xt.", "sad"},
		{":-):-(", "unsure"},
	}

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve10769(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
