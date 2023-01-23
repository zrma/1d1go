package p10700

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve10757(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10757")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			"9223372036854775807 9223372036854775808",
			"18446744073709551615",
		},
		{
			"1 1",
			"2",
		},
		{
			"9 9",
			"18",
		},
		{
			"9 10",
			"19",
		},
		{
			"10 9",
			"19",
		},
		{
			"99999999999999999999999999999999999999999999999 99999999999999999999999999999999999999999999999",
			"199999999999999999999999999999999999999999999998",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve10757(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
