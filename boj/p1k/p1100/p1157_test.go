package p1100

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestSolve1157(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1157")

	//goland:noinspection SpellCheckingInspection
	for _, tt := range []struct {
		s    string
		want string
	}{
		{"Mississipi", "?"},
		{"zZa", "Z"},
		{"z", "Z"},
		{"zZ", "Z"},
		{"baaa", "A"},
		{"cddec", "?"},
		{`AA
BB`, "A"},
		{`AB
BB`, "?"},
		{"AB\nBB", "?"},
		{"AA\nBB", "A"},
	} {
		t.Run(tt.s, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.s))
			writer := utils.NewStringWriter()

			Solve1157(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
