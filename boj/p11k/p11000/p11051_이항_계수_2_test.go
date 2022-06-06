package p11000_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11000"
	"1d1go/utils"
)

func TestSolve11051(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11051")

	for _, tt := range []struct {
		s    string
		want string
	}{
		{"1 0", "1"},
		{"1 1", "1"},
		{"1 2", "0"},
		{"5 2", "10"},
		{"6 1", "6"},
		{"6 2", "15"},
		{"6 3", "20"},
		{"1000 0", "1"},
		{"1000 1", "1000"},
		{"1000 500", "5418"},
	} {
		t.Run(tt.s, func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.s)
			writer := utils.NewStringWriter()

			assert.Eventually(t, func() bool {
				p11000.Solve11051(scanner, writer)
				return true
			}, time.Second, time.Millisecond*100, "시간 초과")

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
