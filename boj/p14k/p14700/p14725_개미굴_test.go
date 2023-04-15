package p14700

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve14725(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/14725")

	tests := []struct {
		give string
		want string
	}{
		{
			`3
2 B A
4 A B C D
2 A C`,
			`A
--B
----C
------D
--C
B
--A
`,
		},
		{
			`4
2 KIWI BANANA
2 KIWI APPLE
2 APPLE APPLE
3 APPLE BANANA KIWI`,
			`APPLE
--APPLE
--BANANA
----KIWI
KIWI
--APPLE
--BANANA
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve14725(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
