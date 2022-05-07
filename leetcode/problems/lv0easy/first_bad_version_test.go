package lv0easy

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveFirstBadVersion(t *testing.T) {
	t.Log("https://leetcode.com/problems/first-bad-version/")

	for _, tt := range []struct {
		maxVer int
		badVer int
	}{
		{5, 4},
		{10, 3},
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
		t.Run(fmt.Sprintf("%d", tt.maxVer), func(t *testing.T) {
			wantCount := int(math.Ceil(math.Log2(float64(tt.maxVer))))
			badVer, callCount := solveFirstBadVersion(tt.maxVer, tt.badVer)
			assert.Equal(t, tt.badVer, badVer)
			assert.LessOrEqual(t, callCount, wantCount, "time complexity limitation")
		})
	}
}
