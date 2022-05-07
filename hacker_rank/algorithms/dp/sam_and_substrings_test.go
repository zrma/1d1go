package dp

import (
	"bufio"
	"encoding/csv"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubStrings(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/sam-and-substrings/problem")

	extremelyLongStr := readExtremelyLongTestData(t)

	for _, tt := range []struct {
		description string
		s           string
		want        int32
	}{
		{
			description: "1 + 6 + 16",
			s:           "16",
			want:        23,
		},
		{
			description: "123 + 12 + 23 + 1 + 2 + 3",
			s:           "123",
			want:        164,
		},
		{
			description: "1212",
			s:           "1212",
			want:        1 + 2 + 1 + 2 + 12 + 21 + 12 + 121 + 212 + 1212,
		},
		{
			description: "1",
			s:           "1",
			want:        1,
		},
		{
			description: "12",
			s:           "12",
			want:        1 + 2 + 12,
		},
		{
			description: "123",
			s:           "123",
			want:        1 + 2 + 12 + 3 + 23 + 123,
		},
		{
			description: "1234",
			s:           "1234",
			want:        1 + 2 + 12 + 3 + 23 + 123 + 4 + 34 + 234 + 1234,
		},
		{
			description: "12345",
			s:           "12345",
			want:        1 + 2 + 12 + 3 + 23 + 123 + 4 + 34 + 234 + 1234 + 5 + 45 + 345 + 2345 + 12345,
		},
		{
			description: "11111",
			s:           "11111",
			want:        1 + 1 + 11 + 1 + 11 + 111 + 1 + 11 + 111 + 1111 + 1 + 11 + 111 + 1111 + 11111,
		},
		{
			description: "a bit long",
			s:           "972698438521",
			want:        445677619,
		},
		{
			description: "extremely long",
			s:           extremelyLongStr,
			want:        597988838,
		},
		{
			description: "previous exception handling",
			s:           "invalid",
			want:        0,
		},
		{
			description: "intermediate exception handling",
			s:           "1234a",
			want:        substrings("12340"),
		},
		{
			description: "intermediate exception handling explanation",
			s:           "1234a",
			want: substrings("1234")*10 +
				1 + 2 + 3 + 4 +
				2 + 3 + 4 +
				3 + 4 +
				4,
		},
	} {
		t.Run(tt.description, func(t *testing.T) {
			got := substrings(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func readExtremelyLongTestData(t *testing.T) string {
	file, err := os.Open("./test_data/sam_and_substrings.csv")
	assert.NoError(t, err)
	defer func() {
		err := file.Close()
		assert.NoError(t, err)
	}()

	r := csv.NewReader(bufio.NewReader(file))
	rows, err := r.ReadAll()
	assert.NoError(t, err)

	assert.Len(t, rows, 1)
	assert.Len(t, rows[0], 1)
	return rows[0][0]
}
