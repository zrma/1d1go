package p2700

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test2744(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2744")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want string
	}{
		{"WrongAnswer", "wRONGaNSWER"},
		{"wRONGaNSWER", "WrongAnswer"},
		{"Hello", "hELLO"},
		{"HelloWorld", "hELLOwORLD"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2744(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
