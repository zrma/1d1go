package p12000_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p12k/p12000"
	"1d1go/utils"
)

func TestSolve12015(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/12015")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`6
10 20 10 30 20 50`,
			"4",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p12000.Solve12015(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
