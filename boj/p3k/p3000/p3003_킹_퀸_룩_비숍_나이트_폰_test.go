package p3000_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p3k/p3000"
	"1d1go/utils"
)

func TestSolve3003(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/3003")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"0 1 2 2 2 7", "1 0 0 0 0 1"},
		{"2 1 2 1 2 1", "-1 0 0 1 0 7"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p3000.Solve3003(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
