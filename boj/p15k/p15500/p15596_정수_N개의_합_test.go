package p15500

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve15596(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/15596")

	const (
		give = `10
1 10 4 9 2 3 8 5 7 6`
		want = "55"
	)

	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve15596(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}

func TestSum(t *testing.T) {
	for i, tt := range []struct {
		give []int
		want int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			15,
		},
		{
			[]int{5, 4, 3, 2},
			14,
		},
		{
			[]int{3, 6, 9},
			18,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := sum(tt.give)
			assert.Equal(t, tt.want, got)
		})
	}
}
