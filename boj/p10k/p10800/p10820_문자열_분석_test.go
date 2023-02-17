package p10800

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve10820(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10820")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			"This is String\nSPACE    1    SPACE\n S a M p L e I n P u T     \n0L1A2S3T4L5I6N7E8",
			`10 2 0 2
0 10 1 8
5 6 0 16
0 8 9 0
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve10820(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
