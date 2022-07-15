package p11800_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"1d1go/boj/p11k/p11800"
	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestSolve11866(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11866")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			"7 3",
			"<3, 6, 2, 7, 5, 1, 4>",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p11800.Solve11866(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
