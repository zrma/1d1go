package lv0easy

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveGuessNumber(t *testing.T) {
	t.Log("https://leetcode.com/problems/guess-number-higher-or-lower/")

	for _, tt := range []struct {
		maxNum int
		target int
	}{
		{5, 4},
		{10, 3},
		{10, 6},
		{100, 37},
		{100, 38},
		{100, 49},
		{100, 50},
		{100, 51},
		{1024, 511},
		{1024, 512},
		{1024, 513},
		{1024, 1024},
	} {
		t.Run(fmt.Sprintf("%d,", tt.maxNum), func(t *testing.T) {
			// 3진 탐색, 1회 탐색에 2번씩 호출
			wantCount := int(math.Ceil((math.Log(float64(tt.maxNum)))/math.Log(3))*2) + 1
			got, callCount := solveGuessNumber(tt.maxNum, tt.target)
			assert.Equal(t, tt.target, got)
			assert.LessOrEqual(t, callCount, wantCount, "성능 제한")
		})
	}

}
func TestGuessNumber(t *testing.T) {
	assert.NotPanics(t, func() {
		const want = -1
		got := guessNumber(0)
		assert.Equal(t, want, got)
	})
}
