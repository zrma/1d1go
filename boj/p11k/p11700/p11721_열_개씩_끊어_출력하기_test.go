package p11700

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve11721(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11721")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want string
	}{
		{
			"BaekjoonOnlineJudge",
			`BaekjoonOn
lineJudge
`,
		},
		{
			"OneTwoThreeFourFiveSixSevenEightNineTen",
			`OneTwoThre
eFourFiveS
ixSevenEig
htNineTen
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve11721(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
