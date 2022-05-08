package practice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportResult(t *testing.T) {
	t.Log("https://programmers.co.kr/learn/courses/30/lessons/92334")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		idList []string
		report []string
		k      int
		want   []int
	}{
		{
			[]string{"muzi", "frodo", "apeach", "neo"},
			[]string{"muzi frodo", "apeach frodo", "frodo neo", "muzi neo", "apeach muzi"},
			2,
			[]int{2, 1, 1, 0},
		},
		{
			[]string{"con", "ryan"},
			[]string{"ryan con", "ryan con", "ryan con", "ryan con"},
			3,
			[]int{0, 0},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := reportResult(tt.idList, tt.report, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}
