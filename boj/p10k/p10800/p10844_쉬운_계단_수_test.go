package p10800_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10800"
	"1d1go/utils"
)

func TestSolve10844(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10844")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"1", "9"},
		{"2", "17"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p10800.Solve10844(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
