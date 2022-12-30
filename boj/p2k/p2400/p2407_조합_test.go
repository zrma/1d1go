package p2400_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2400"
)

func TestSolve2407(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2407")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"1 0", "1"},
		{"1 1", "1"},
		{"1 2", "0"},
		{"5 2", "10"},
		{"6 1", "6"},
		{"6 2", "15"},
		{"6 3", "20"},
		{`100 6`, "1192052400"},
		{"1000 0", "1"},
		{"1000 1", "1000"},
		{"1000 500", "270288240945436569515614693625975275496152008446548287007392875106625428705522193898612483924502370165362606085021546104802209750050679917549894219699518475423665484263751733356162464079737887344364574161119497604571044985756287880514600994219426752366915856603136862602484428109296905863799821216320"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p2400.Solve2407(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
