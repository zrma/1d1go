package p5500

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve5524(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/5524")

	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		give string
		want string
	}{
		{
			`3
WatanabE
ITO
YamaMoto`,
			`watanabe
ito
yamamoto
`,
		},
		{
			`4
SUZUKI
tanaka
tAkAhAshi
SuZuKi`,
			`suzuki
tanaka
takahashi
suzuki
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve5524(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
