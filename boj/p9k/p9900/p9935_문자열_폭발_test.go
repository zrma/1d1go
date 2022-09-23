package p9900_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p9k/p9900"
	"1d1go/utils"
)

func TestSolve9935(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/9935")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`mirkovC4nizCC44
C4`,
			`mirkovniz`,
		},
		{
			`12ab112ab2ab
12ab`,
			`FRULA`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p9900.Solve9935(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
