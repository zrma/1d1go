package p2500

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2525(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2525")

	for i, tt := range []struct {
		hour, minute, duration int
		wantHour, wantMinute   int
	}{
		{14, 30, 20, 14, 50},
		{17, 40, 80, 19, 0},
		{23, 48, 25, 0, 13},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			gotHour, gotMinute := Solve2525(tt.hour, tt.minute, tt.duration)
			assert.Equal(t, tt.wantHour, gotHour)
			assert.Equal(t, tt.wantMinute, gotMinute)
		})
	}
}
