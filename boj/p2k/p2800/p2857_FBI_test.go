package p2800

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2857(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2857")

	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		give string
		want string
	}{
		{
			`N-FBI1
9A-USKOK
I-NTERPOL
G-MI6
RF-KGB1`,
			"1 ",
		},
		{
			`N321-CIA
F3-B12I
F-BI-12
OVO-JE-CIA
KRIJUMCAR1`,
			"HE GOT AWAY!",
		},
		{
			`47-FBI
BOND-007
RF-FBI18
MARICA-13
13A-FBILL`,
			"1 3 5 ",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2857(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
