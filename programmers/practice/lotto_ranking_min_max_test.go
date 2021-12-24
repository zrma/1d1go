package practice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLottoRankingMinMax(t *testing.T) {
	t.Log("https://programmers.co.kr/learn/courses/30/lessons/77484?language=go")

	for i, tt := range []struct {
		lottoNumbers []int
		winNumbers   []int
		want         []int
	}{
		{
			[]int{44, 1, 0, 0, 31, 25},
			[]int{31, 10, 45, 1, 6, 19},
			[]int{3, 5},
		},
		{
			[]int{0, 0, 0, 0, 0, 0},
			[]int{3, 8, 19, 20, 40, 15, 25},
			[]int{1, 6},
		},
		{
			[]int{45, 4, 35, 20, 3, 9},
			[]int{20, 9, 3, 45, 4, 35},
			[]int{1, 1},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := lottoRankingMinMax(tt.lottoNumbers, tt.winNumbers)
			assert.Equal(t, tt.want, got)
		})
	}
}
