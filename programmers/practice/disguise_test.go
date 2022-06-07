package practice_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/programmers/practice"
)

func TestDisguise(t *testing.T) {
	t.Log("https://programmers.co.kr/learn/courses/30/lessons/42578?language=go")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		clothes [][]string
		want    int
	}{
		{
			clothes: [][]string{
				{"yellowhat", "headgear"},
				{"bluesunglasses", "eyewear"},
				{"green_turban", "headgear"},
			},
			want: 5,
		},
		{
			clothes: [][]string{
				{"crowmask", "face"},
				{"bluesunglasses", "face"},
				{"smoky_makeup", "face"},
			},
			want: 3,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := practice.Disguise(tt.clothes)
			assert.Equal(t, tt.want, got)
		})
	}
}
