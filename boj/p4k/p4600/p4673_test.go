package p4600_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p4k/p4600"
	"1d1go/utils"
)

func TestSolve4673(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/4673")

	for _, tt := range []struct {
		n    int
		want []int
	}{
		{50, []int{1, 3, 5, 7, 9, 20, 31, 42}},
		{99, []int{1, 3, 5, 7, 9, 20, 31, 42, 53, 64, 75, 86, 97}},
	} {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			want := make([]string, len(tt.want))
			for i, v := range tt.want {
				want[i] = strconv.Itoa(v)
			}
			got, err := utils.GetPrinted(func() {
				p4600.Solve4673(tt.n)
			})
			assert.NoError(t, err)
			assert.Equal(t, want, got)
		})
	}
}
