package p12600

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve12605(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/12605")

	tests := []struct {
		give string
		want string
	}{
		{
			`3
this is a test
foobar
all your base`,
			`Case #1: test a is this
Case #2: foobar
Case #3: base your all
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve12605(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
