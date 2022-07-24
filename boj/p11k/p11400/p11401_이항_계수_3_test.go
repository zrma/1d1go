package p11400_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11400"
	"1d1go/utils"
)

func TestSolve11401(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11401")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"5 2", "10"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p11400.Solve11401(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
