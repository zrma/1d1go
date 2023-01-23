package p1100

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1152(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1152")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"The Curious Case of Benjamin Button", "6"},
		{" The first character is a blank", "6"},
		{"The last character is a blank  ", "6"},
		{" ", "0"},
		{"Word ", "1"},
		{" Word", "1"},
		{`These words are
edge cases`, "3"},
		{"These words are\n edge cases", "3"},
		{"These words are \n edge cases", "3"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1152(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
