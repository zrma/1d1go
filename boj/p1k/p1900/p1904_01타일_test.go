package p1900

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSolve1904(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1904")

	tests := []struct {
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
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			assert.Eventually(t, func() bool {
				Solve1904(reader, writer)

				err := writer.Flush()
				assert.NoError(t, err)

				got := buf.String()
				return assert.Equal(t, tt.want, got)
			}, time.Second, time.Millisecond*100, "시간 초과")
		})
	}
}
