package p4100_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p4k/p4100"
	"1d1go/utils"
)

func TestSolve4195(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/4195")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`2
3
Fred Barney
Barney Betty
Betty Wilma
3
Fred Barney
Betty Wilma
Barney Betty`,
			`2
3
4
2
2
4
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p4100.Solve4195(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
