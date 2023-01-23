package practice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisguise(t *testing.T) {
	t.Log("https://programmers.co.kr/learn/courses/30/lessons/42578?language=go")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give [][]string
		want int
	}{
		{
			give: [][]string{
				{"yellowhat", "headgear"},
				{"bluesunglasses", "eyewear"},
				{"green_turban", "headgear"},
			},
			want: 5,
		},
		{
			give: [][]string{
				{"crowmask", "face"},
				{"bluesunglasses", "face"},
				{"smoky_makeup", "face"},
			},
			want: 3,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := Disguise(tt.give)
			assert.Equal(t, tt.want, got)
		})
	}
}
