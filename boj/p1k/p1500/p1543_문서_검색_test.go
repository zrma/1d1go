package p1500

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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
			"a a a  \na  ",
			"1",
		},
		{
			`abababc
ababc`,
			"1",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1543(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
