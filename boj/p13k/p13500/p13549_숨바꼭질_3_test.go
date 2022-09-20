package p13300_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	p13300 "1d1go/boj/p13k/p13500"
	"1d1go/utils"
)

func TestSolve13549(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/13549")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"5 17", "2"},
		{"1 2", "0"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p13300.Solve13549(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
