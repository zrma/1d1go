package p17600_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p17k/p17600"
)

func TestSolve17609(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/17609")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`7
abba
summuus
xabba
xabbay
comcom
comwwmoc
comwwtmoc`,
			`0
1
1
2
2
0
1
`,
		},
		{
			`1
aacadaaaa`,
			`2
`,
		},
		{
			`21
abbab
aab
aaab
aaaab
aaaaab
aaaaaab
axaaxaa
abcddadca
aabcbcaa
ababbabaa
abca
babba
sumumuus
XYXYAAYXY
abc
cccfccfcc
abcddcdba
ppbpppb
aabcdeddcba
aabab
aapqbcbqpqaa`,
			`1
1
1
1
1
1
1
2
1
1
1
1
1
1
2
1
1
2
2
2
1
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p17600.Solve17609(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
