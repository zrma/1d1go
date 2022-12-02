package p4200_test

import (
	"1d1go/boj/p4k/p4200"
	"1d1go/utils"
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestSolve4299(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/4299")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"3 1", "2 1"},
		{"3 2", "-1"},
		{"3 3", "3 0"},
		{"3 4", "-1"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p4200.Solve4299(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
