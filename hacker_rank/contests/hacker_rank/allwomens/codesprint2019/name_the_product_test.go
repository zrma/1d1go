package codesprint2019

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductName(t *testing.T) {
	t.Log("https://www.hackerrank.com/contests/hackerrank-all-womens-codesprint-2019/challenges/name-the-product")

	//noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give []string
		want string
	}{
		{
			give: []string{"bubby", "bunny", "berry"},
			want: "zzzzz",
		},
		{
			give: []string{"ready", "stedy", "zebra"},
			want: "yzzzz",
		},
		{
			give: []string{
				"aaaaa", "bbbbb", "ccccc", "ddddd", "eeeee", "fffff", "ggggg",
				"hhhhh", "iiiii", "jjjjj", "kkkkk", "lllll", "mmmmm", "nnnnn",
				"ooooo", "ppppp", "qqqqq", "rrrrr", "sssss", "ttttt", "uuuuu",
				"vvvvv", "wwwww", "xxxxx", "yyyyy", "zzzzz", "zzzzz",
			},
			want: "yyyyy",
		},
		{
			give: []string{"uuuuu", "vvvvv", "wwwww", "yyyyy", "zzzzz"},
			want: "xxxxx",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := productName(tt.give)
			assert.Equal(t, tt.want, got)
		})
	}
}
