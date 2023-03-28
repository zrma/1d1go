package p1300

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1371(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1371")

	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		give string
		want string
	}{
		{
			`english is a west germanic
language originating in england
and is the first language for
most people in the united
kingdom the united states
canada australia new zealand
ireland and the anglophone
caribbean it is used
extensively as a second
language and as an official
language throughout the world
especially in common wealth
countries and in many
international organizations`,
			"a",
		},
		{
			`baekjoon online judge`,
			"eno",
		},
		{
			`abc a`,
			"a",
		},
		{
			`abc
ab`,
			"ab",
		},
		{
			`amanda forsaken bloomer meditated gauging knolls
betas neurons integrative expender commonalities
latins antidotes crutched bandwidths begetting
prompting dog association athenians christian ires
pompousness percolating figured bagatelles bursted
ninth boyfriends longingly muddlers prudence puns
groove deliberators charter collectively yorks
daringly antithesis inaptness aerosol carolinas
payoffs chumps chirps gentler inexpressive morales`,
			"e",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1371(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
