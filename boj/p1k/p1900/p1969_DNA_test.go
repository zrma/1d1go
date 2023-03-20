package p1900

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1969(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1969")

	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		give string
		want string
	}{
		{
			`5 8
TATGATAC
TAAGCTAC
AAAGATCC
TGAGATAC
TAAGATGT`,
			`TAAGATAC
7`,
		},
		{
			`4 10
ACGTACGTAC
CCGTACGTAG
GCGTACGTAT
TCGTACGTAA`,
			`ACGTACGTAA
6`,
		},
		{
			`6 10
ATGTTACCAT
AAGTTACGAT
AACAAAGCAA
AAGTTACCTT
AAGTTACCAA
TACTTACCAA`,
			`AAGTTACCAA
12`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1969(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
