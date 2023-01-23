package p4900

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve4949(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/4949")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`So when I die (the [first] I will see in (heaven) is a score list).
[ first in ] ( first out ).
Half Moon tonight (At least it is better than no Moon at all].
A rope may form )( a trail in a maze.
Help( I[m being held prisoner in a fortune cookie factory)].
([ (([( [ ] ) ( ) (( ))] )) ]).
 .
.`,
			`yes
yes
no
no
no
yes
yes
`,
		},
		{
			`Help( I[m being held prisoner in a fortune cookie factory)].
.`,
			`no
`,
		},
		{
			`Help( I[m being held prisoner in a fortune cookie factory)].
.


`,
			`no
`,
		},
		{
			`].
.`,
			`no
`,
		},
		{
			`).
.`,
			`no
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve4949(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
