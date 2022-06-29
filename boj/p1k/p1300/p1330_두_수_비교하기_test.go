package p1300_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1300"
	"1d1go/utils"
)

func TestSolve1330(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1330")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"1 2", "<"},
		{"10 2", ">"},
		{"5 5", "=="},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p1300.Solve1330(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
