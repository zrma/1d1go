package p1500_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1500"
)

func TestSolve1543(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1543")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`ababababa
aba`,
			"2",
		},
		{
			`a a a a a
a a`,
			"2",
		},
		{
			`ababababa
ababa`,
			"1",
		},
		{
			`aaaaaaa
aa`,
			"3",
		},
		{
			`mdm
md`,
			"1",
		},
		{
			`a a a  
a  `,
			"1",
		},
		{
			`abababc
ababc`,
			"1",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewScanner(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p1500.Solve1543(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
