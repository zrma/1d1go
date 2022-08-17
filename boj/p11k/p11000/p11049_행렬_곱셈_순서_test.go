package p11000_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11000"
	"1d1go/utils"
)

func TestSolve11049(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11049")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3
5 3
3 2
2 6`,
			`90`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p11000.Solve11049(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
