package p4600

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

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
				Solve4673(tt.n)
			})
			assert.NoError(t, err)
			assert.Equal(t, want, got)
		})
	}
}

func TestD(t *testing.T) {
	for _, tt := range []struct {
		n    int
		want int
	}{
		{1, 2},
		{33, 39},
		{39, 51},
		{51, 57},
		{57, 69},
		{69, 84},
		{84, 96},
		{96, 111},
		{111, 114},
		{114, 120},
		{120, 123},
	} {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			got := d(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
