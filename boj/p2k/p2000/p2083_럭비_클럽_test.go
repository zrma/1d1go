package p2000_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2000"
	"1d1go/utils"
)

func TestSolve2083(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2083")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`Joe 16 34
Bill 18 65
Billy 17 65
Sam 17 85
# 0 0`,
			`Joe Junior
Bill Senior
Billy Junior
Sam Senior
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p2000.Solve2083(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
