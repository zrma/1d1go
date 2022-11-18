package p11700_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11700"
	"1d1go/utils"
)

func TestSolve11758(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11758")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`1 1
5 5
7 3`,
			"-1",
		},
		{
			`1 1
3 3
5 5`,
			"0",
		},
		{
			`1 1
7 3
5 5`,
			"1",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p11700.Solve11758(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
