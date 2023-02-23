package p1300

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1357(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1357")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			"123 100",
			"223",
		},
		{
			"111 111",
			"222",
		},
		{
			"5 5",
			"1",
		},
		{
			"1000 1",
			"2",
		},
		{
			"456 789",
			"1461",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1357(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestReverse(t *testing.T) {
	for i, tt := range []struct {
		give int
		want int
	}{
		{123, 321},
		{100, 1},
		{1, 1},
		{0, 0},
		{1000, 1},
		{456, 654},
		{101, 101},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := reverse(tt.give)
			assert.Equal(t, tt.want, got)
		})
	}
}
