package p18100_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p18k/p18100"
)

func TestSolve18108(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/18108")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"2541", "1998"},
		{"543", "0"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p18100.Solve18108(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
