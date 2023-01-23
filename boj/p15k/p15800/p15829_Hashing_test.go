package p15800

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve15829(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/15829")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5
abcde`,
			"4739715",
		},
		{
			`3
zzz`,
			"25818",
		},
		{
			`1
i`,
			"9",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve15829(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
