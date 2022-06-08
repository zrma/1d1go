package p2000_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2000"
	"1d1go/utils"
)

func TestSolve2004(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2004")

	for _, tt := range []struct {
		s    string
		want string
	}{
		{"25 12", "2"},
		{"5 1", "0"},
		{"5 4", "0"},
		{"1 1", "0"},
		{"10 2", "0"},
		{"2000000000 1000000000", "1"},
		{"2000000000 1", "9"},
	} {
		t.Run(tt.s, func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.s)
			writer := utils.NewStringWriter()

			p2000.Solve2004(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSolve2004_Performance(t *testing.T) {
	const (
		s    = "2000000000 1"
		want = "9"
	)

	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	assert.Eventually(t, func() bool {
		p2000.Solve2004(scanner, writer)

		err := writer.Flush()
		assert.NoError(t, err)

		got := writer.String()
		return assert.Equal(t, want, got)
	}, time.Second, time.Millisecond*100, "시간 초과")
}
