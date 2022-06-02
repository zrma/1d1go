package p2900

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestSolve2941(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2941")

	//goland:noinspection SpellCheckingInspection
	for _, tt := range []struct {
		s    string
		want string
	}{
		{"c", "1"},
		{"c=", "1"},
		{"c-", "1"},
		{"d", "1"},
		{"dz", "2"},
		{"dz=", "1"},
		{"d-", "1"},
		{"l", "1"},
		{"lj", "1"},
		{"n", "1"},
		{"nj", "1"},
		{"s", "1"},
		{"s=", "1"},
		{"z", "1"},
		{"z=", "1"},
		{"ljes=njak", "6"},
		{"ddz=z=", "3"},
		{"dz=z=", "2"},
		{"nljj", "3"},
		{"c=c=", "2"},
		{"dz=ak", "3"},
		{"", "0"},
	} {
		t.Run(tt.s, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.s))
			writer := utils.NewStringWriter()

			Solve2941(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
