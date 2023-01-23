package p11400

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve11478(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11478")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want string
	}{
		{"ababc", "12"},
		{"abcde", "15"},
		{"abcdefghij", "55"},
		{"abcdefghab", "52"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve11478(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
