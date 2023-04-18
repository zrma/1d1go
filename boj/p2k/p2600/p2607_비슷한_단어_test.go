package p2600

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2607(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2607")

	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		give string
		want string
	}{
		{
			`4
DOG
GOD
GOOD
DOLL`,
			`2`,
		},
		{
			`7
DOG
GOD
GOOD
DOLL
DATE
Q
DOLLY`,
			`2`,
		},
		{
			`2
ABCD
DCBA`,
			`1`,
		},
		{
			`2
ABCD
ABDD`,
			`1`,
		},
		{
			`2
ABCD
ABDDD`,
			`0`,
		},
		{
			`2
ABCD
ABDEC`,
			`1`,
		},
		{
			`5
DOG
DO
GOD
DOH
DDOG`,
			`4`,
		},
		{
			`10
ABABA
BBAAC
BBAAD
BBAAE
BBAAB
ACQAC
LKJDS
OKWLN
OKNLS
PONLM`,
			`4`,
		},
		{
			`10
AAAB
ABBA
ABBB
AAABC
AAABB
AABBB
PABB
ABCD
AB
ABA`,
			`4`,
		},
		{
			`4
ABC
ABCD
BC
PCD`,
			`2`,
		},
		{
			`8
DOG
GOD
GOOD
DOGL
GD
PGG
G
DOGOD`,
			`4`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2607(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
