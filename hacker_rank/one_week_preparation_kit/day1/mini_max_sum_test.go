package day1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestMiniMaxSum(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/mini-max-sum/problem")
	t.Log("https://www.hackerrank.com/challenges/one-week-preparation-kit-mini-max-sum/problem")

	for i, tt := range []struct {
		given []int32
		want  []string
	}{
		{
			given: []int32{1, 2, 3, 4, 5},
			want:  []string{"10 14"},
		},
		{
			given: []int32{1, 3, 5, 7, 9},
			want:  []string{"16 24"},
		},
	} {
		t.Run(fmt.Sprintf(" %d", i), func(t *testing.T) {
			got, err := utils.GetPrinted(func() {
				miniMaxSum(tt.given)
			})
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
