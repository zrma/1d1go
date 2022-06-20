package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeConversion(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/time-conversion/problem")
	t.Log("https://www.hackerrank.com/challenges/one-week-preparation-kit-time-conversion/problem")

	for _, tt := range []struct {
		give string
		want string
	}{
		{"07:05:45PM", "19:05:45"},
		{"12:01:00PM", "12:01:00"},
		{"12:01:00AM", "00:01:00"},
		{"11:59:59AM", "11:59:59"},
		{"11:59:59PM", "23:59:59"},
	} {
		t.Run(tt.give, func(t *testing.T) {
			assert.NotPanics(t, func() {
				got := timeConversion(tt.give)
				assert.Equal(t, tt.want, got)
			})
		})
	}
}
