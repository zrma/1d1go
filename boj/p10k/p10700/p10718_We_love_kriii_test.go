package p10700

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve10718(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10718")

	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		want string
	}{
		{
			`강한친구 대한육군
강한친구 대한육군
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)
			Solve10718(writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
