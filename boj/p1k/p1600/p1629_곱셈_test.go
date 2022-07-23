package p1600_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1600"
	"1d1go/utils"
)

func TestSolve1629(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1629")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"10 11 12", "4"},
		{"10 0 12", "1"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p1600.Solve1629(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
