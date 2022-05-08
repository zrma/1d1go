package ps_basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumCost(t *testing.T) {
	t.Log("https://www.hackerrank.com/skills-verification/problem_solving_basic")

	for i, tt := range []struct {
		cost       []int32
		labels     []string
		dailyCount int32
		want       int32
	}{
		{
			cost:       []int32{2, 5, 3, 11, 1},
			labels:     []string{"legal", "illegal", "legal", "illegal", "legal"},
			dailyCount: 2,
			want:       10,
		},
		{
			cost:       []int32{2},
			labels:     []string{"illegal"},
			dailyCount: 1,
			want:       0,
		},
		{
			cost:       []int32{0, 3, 2, 3, 4},
			labels:     []string{"legal", "legal", "illegal", "legal", "legal"},
			dailyCount: 1,
			want:       5,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := maxCost(tt.cost, tt.labels, tt.dailyCount)
			assert.Equal(t, tt.want, got)
		})
	}
}
