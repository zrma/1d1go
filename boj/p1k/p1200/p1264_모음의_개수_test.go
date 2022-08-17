package p1200_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1200"
	"1d1go/utils"
)

func TestSolve1264(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1264")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`How are you today?
Quite well, thank you, how about yourself?
I live at number twenty four.
#`,
			`7
14
9
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p1200.Solve1264(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
