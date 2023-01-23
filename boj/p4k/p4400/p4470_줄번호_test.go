package p4400

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2393(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2393")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5
Lionel Cosgrove
Alice
Columbus and Tallahassee
Shaun and Ed
Fido`,
			`1. Lionel Cosgrove
2. Alice
3. Columbus and Tallahassee
4. Shaun and Ed
5. Fido
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve4470(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
