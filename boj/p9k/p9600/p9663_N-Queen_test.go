package p9600_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p9k/p9600"
	"1d1go/utils"
)

func TestSolve9663(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/9663")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"1", "1"},
		{"2", "0"},
		{"3", "0"},
		{"4", "2"},
		{"5", "10"},
		{"6", "4"},
		{"7", "40"},
		{"8", "92"},
		{"9", "352"},
		{"10", "724"},
		{"11", "2680"},
		{"12", "14200"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p9600.Solve9663(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
