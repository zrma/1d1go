package p11600_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11600"
	"1d1go/utils"
)

func TestSolve11653(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11653")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"1", ""},
		{
			"72",
			`2
2
2
3
3
`,
		},
		{
			"3",
			`3
`,
		},
		{
			"6",
			`2
3
`,
		},
		{
			"2",
			`2
`,
		},
		{
			"9",
			`3
3
`,
		},
		{
			"9991",
			`97
103
`,
		},
		{
			"379721",
			`379721
`,
		},
		{
			"123456789",
			`3
3
3607
3803
`,
		},
		{
			"987654321",
			`3
3
17
17
379721
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p11600.Solve11653(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
