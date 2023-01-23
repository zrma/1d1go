package p17300

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve17386(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/17386")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`1 1 5 5
1 5 5 1`,
			"1",
		},
		{
			`1 1 5 5
6 10 10 6`,
			"0",
		},
		{
			`0 0 1 1
1 1 2 2`,
			"1",
		},
		{
			`1 1 0 0
1 1 2 2`,
			"1",
		},
		{
			`0 0 1 1
1 1 2 2`,
			"1",
		},
		{
			`0 0 1 1
2 2 3 3`,
			"0",
		},
		{
			`2 2 1 1
0 0 1 1`,
			"1",
		},
		{
			`3 3 2 2
0 0 1 1`,
			"0",
		},
		{
			`0 0 2 2
1 1 3 3`,
			"1",
		},
		{
			`1 1 5 5
3 3 1 3`,
			"1",
		},
		{
			`0 0 1 1
2 2 3 3`,
			"0",
		},
		{
			`0 0 1 1
1 1 2 2`,
			"1",
		},
		{
			`0 0 1 1
1 2 2 2`,
			"0",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve17386(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPoint_LessThan(t *testing.T) {
	for i, tt := range []struct {
		give []Point
		want bool
	}{
		{
			[]Point{{1, 1}, {5, 5}},
			true,
		},
		{
			[]Point{{5, 5}, {1, 1}},
			false,
		},
		{
			[]Point{{1, 1}, {1, 1}},
			false,
		},
		{
			[]Point{{1, 1}, {1, 2}},
			true,
		},
		{
			[]Point{{1, 2}, {1, 1}},
			false,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := tt.give[0].LessThan(tt.give[1])
			assert.Equal(t, tt.want, got)
		})
	}
}
