package p11000_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"1d1go/boj/p11k/p11000"
	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestSolve11047(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11047")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`10 4200
1
5
10
50
100
500
1000
5000
10000
50000`,
			"6",
		},
		{
			`10 4790
1
5
10
50
100
500
1000
5000
10000
50000`,
			"12",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p11000.Solve11047(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
