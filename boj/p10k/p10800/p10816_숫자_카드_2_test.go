package p10800

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve10816(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10816")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`10
6 3 2 10 10 10 -10 -10 7 3
8
10 9 -5 2 3 4 5 -10`,
			"3 0 0 1 2 0 0 2 ",
		},
	} {
		t.Run(fmt.Sprintf("%d/map", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve10816WithMap(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})

		t.Run(fmt.Sprintf("%d/bin search", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve10816(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
