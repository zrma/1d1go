package p2900_test

import (
	_ "embed"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2900"
	"1d1go/utils"
)

func TestSolve2981(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2981")

	for i, tt := range []struct {
		s    string
		want string
	}{
		{
			`3
6
34
38`,
			"2 4 ",
		},
		{
			`3
1
17
81`,
			"2 4 8 16 ",
		},
		{
			`5
5
17
23
14
83`,
			"3 ",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.s)
			writer := utils.NewStringWriter()

			p2900.Solve2981(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

//go:embed test_data/p2981_give.txt
var p2981give string

//go:embed test_data/p2981_want.txt
var p2981want string

func TestSolve2981_Performance(t *testing.T) {
	scanner := utils.NewStringScanner(p2981give)
	writer := utils.NewStringWriter()

	assert.Eventually(t, func() bool {
		p2900.Solve2981(scanner, writer)
		return true
	}, time.Second, time.Millisecond*100, "시간 초과")

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	got = strings.TrimRight(got, " ") + "\n"
	assert.Equal(t, p2981want, got)
}
