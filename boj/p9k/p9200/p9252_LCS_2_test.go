package p9200

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve9252(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/9252")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`ACAYKP
CAPCAK`,
			`4
ACAK`,
		},
		{
			`ABCDEF
GBCDEF`,
			`5
BCDEF`,
		},
		{
			`ABCDEF
GBCDFE`,
			`4
BCDF`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve9252(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
