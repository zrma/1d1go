package p11000_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"1d1go/boj/p11k/p11000"
	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestSolve11054(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11054")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`10
1 5 2 1 4 3 4 5 2 1`,
			"7",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p11000.Solve11054(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
