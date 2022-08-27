package p1900_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1900"
	"1d1go/utils"
)

func TestSolve1904(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1904")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"0", "1"},
		{"1", "1"},
		{"2", "2"},
		{"3", "3"},
		{"4", "5"},
		{"5", "8"},
		{"1000000", "7871"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			assert.Eventually(t, func() bool {
				p1900.Solve1904(reader, writer)

				err := writer.Flush()
				assert.NoError(t, err)

				got := writer.String()
				return assert.Equal(t, tt.want, got)
			}, time.Second, time.Millisecond*100, "시간 초과")
		})
	}
}
