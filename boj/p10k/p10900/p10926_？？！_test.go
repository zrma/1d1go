package p10900_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10900"
)

func TestSolve10926(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10926")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want string
	}{
		{"joonas", "joonas??!"},
		{"baekjoon", "baekjoon??!"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p10900.Solve10926(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
