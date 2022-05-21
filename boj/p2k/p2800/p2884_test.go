package p2800_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2800"
)

func TestSolve2884(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2884")

	for i, tt := range []struct {
		hour, minute         int
		wantHour, wantMinute int
	}{
		{10, 10, 9, 25},
		{0, 30, 23, 45},
		{23, 40, 22, 55},
		{0, 45, 0, 0},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			gotHour, gotMinute := p2800.Solve2884(tt.hour, tt.minute)
			assert.Equal(t, tt.wantHour, gotHour)
			assert.Equal(t, tt.wantMinute, gotMinute)
		})
	}
}
