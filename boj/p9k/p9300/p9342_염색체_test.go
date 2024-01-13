package p9300

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve9342(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/9342")

	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		give string
		want string
	}{
		{
			`15
AFC
AAFC
AAAFFCC
AAFCC
BAFC
QWEDFGHJMNB
DFAFCB
ABCDEFC
DADC
SDFGHJKLQWERTYU
AAAAAAAAAAAAABBBBBBBBBBBBBBCCCCCCCCCCCCCCCCCCDDDDDDDDDDDEEEEEEEEEEEEEEEFFFFFFFFC
AAAFFFFFBBBBCCCAAAFFFF
ABCDEFAAAFFFCCCABCDEF
AFCP
AAFFCPP`,
			`Infected!
Infected!
Infected!
Infected!
Infected!
Good
Good
Good
Good
Good
Good
Good
Good
Good
Good
`,
		},
		{
			`1

`,
			`Good
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve9342(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIsMatched_Panic(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	const (
		giveStr          = "FooBar"
		wantRegexPattern = `^[ABCDEF]?A+F+C+[ABCDEF]?$`
	)
	regexpMatchString = func(pattern string, s string) (matched bool, err error) {
		assert.Equal(t, wantRegexPattern, pattern)
		assert.Equal(t, giveStr, s)
		return true, fmt.Errorf("test error")
	}
	defer func() { regexpMatchString = regexp.MatchString }()

	got := isMatched(giveStr)
	assert.False(t, got)
}
