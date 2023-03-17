package p11000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve11098(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11098")

	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		give string
		want string
	}{
		{
			`2
3
10 Iversen
1000000 Nannskog
2000000 Ronaldinho
2
1000000 Maradona
999999 Batistuta`,
			`Ronaldinho
Maradona
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve11098(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
