package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestClassAndInstance(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-class-vs-instance/problem")

	for _, tt := range []struct {
		age  int
		want []string
	}{
		{
			age: -1,
			want: []string{
				"Age is not valid, setting age to 0.",
				"You are young.",
				"You are young.",
				"",
			},
		},
		{
			age: 10,
			want: []string{
				"You are young.",
				"You are a teenager.",
				"",
			},
		},
		{
			age: 16,
			want: []string{
				"You are a teenager.",
				"You are old.",
				"",
			},
		},
		{
			age: 18,
			want: []string{
				"You are old.",
				"You are old.",
				"",
			},
		},
	} {
		t.Run(fmt.Sprintf("%d", tt.age), func(t *testing.T) {
			got, err := utils.GetPrinted(func() {
				classAndInstance(tt.age)
			})
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
