package p11700

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve11719(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11719")

	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		give string
		want string
	}{
		{
			"    Hello\n\nBaekjoon     \n   Online Judge    ",
			"    Hello\n\nBaekjoon     \n   Online Judge    \n",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve11719(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
