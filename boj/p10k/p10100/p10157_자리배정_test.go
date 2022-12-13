package p10100_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10100"
)

func TestSolve10157(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10157")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`7 6
11`,
			"6 6",
		},
		{
			`7 6
87`,
			"0",
		},
		{
			`100 100
3000`,
			"9 64",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p10100.Solve10157(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
