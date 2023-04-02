package p11200

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve11283(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11283")

	tests := []struct {
		give string
		want string
	}{
		{"가", "1"},
		{"나", "1177"},
		{"힣", "11172"},
		{"백", "4146"},
		{"준", "7425"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve11283(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
