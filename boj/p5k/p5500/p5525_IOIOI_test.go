package p5500_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p5k/p5500"
	"1d1go/utils"
)

func TestSolve5525(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/5525")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`1
13
OOIOIOIOIIOII`,
			"4",
		},
		{
			`2
13
OOIOIOIOIIOII`,
			"2",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p5500.Solve5525(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
