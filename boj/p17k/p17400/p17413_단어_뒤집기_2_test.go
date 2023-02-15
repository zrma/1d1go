package p17400

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve17413(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/17413")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want string
	}{
		{"baekjoon online judge", "noojkeab enilno egduj"},
		{"<open>tag<close>", "<open>gat<close>"},
		{"<ab cd>ef gh<ij kl>", "<ab cd>fe hg<ij kl>"},
		{"one1 two2 three3 4fourr 5five 6six", "1eno 2owt 3eerht rruof4 evif5 xis6"},
		{"<int><max>2147483647<long long><max>9223372036854775807", "<int><max>7463847412<long long><max>7085774586302733229"},
		{"<problem>17413<is hardest>problem ever<end>", "<problem>31471<is hardest>melborp reve<end>"},
		{"<   space   >space space space<    spa   c e>", "<   space   >ecaps ecaps ecaps<    spa   c e>"},
		{"<abc><abc>ab", "<abc><abc>ba"},
		{"<abc><abc>ab<abc>", "<abc><abc>ba<abc>"},
		{"<abc><abc><abc>", "<abc><abc><abc>"},
		{"abcdef aac", "fedcba caa"},
		{`abcdef aac
`, "fedcba caa"},
		{"abc<def>abc", "cba<def>cba"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve17413(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
