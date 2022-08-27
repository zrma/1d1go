package p5600_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p5k/p5600"
	"1d1go/utils"
)

func TestSolve5622(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/5622")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want string
	}{
		{"WA", "13"},
		{"WA\n", "13"},
		{"WA\r\n", "13"},
		{"UNUCIC", "36"},
		{"C", "3"},
		{"D", "4"},
		{"F", "4"},
		{"G", "5"},
		{"P", "8"},
		{"S", "8"},
		{"T", "9"},
		{"V", "9"},
		{"W", "10"},
		{"Z", "10"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p5600.Solve5622(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
