package p1100

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1157(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1157")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
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
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1157(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
