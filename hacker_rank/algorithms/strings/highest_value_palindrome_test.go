package strings

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHighestValuePalindrome(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/richie-rich/problem")

	extremelyLongGiven, extremelyLongWant := readExtremelyLongTestData(t)

	for i, tt := range []struct {
		s    string
		n, k int32
		want string
	}{
		{"3943", 4, 1, "3993"},
		{"092282", 6, 3, "992299"},
		{"0011", 4, 1, "-1"},
		{"1100", 4, 1, "-1"},
		{"", 0, 0, "-1"},
		{"1", 1, 0, "1"},
		{"1", 1, 1, "9"},
		{"01", 2, 1, "11"},
		{"00", 2, 2, "99"},
		{"102", 3, 2, "909"},
		{"12321", 5, 1, "12921"},
		{"11111111", 8, 4, "99111199"},
		{"11111111", 8, 5, "99111199"},
		{"1111111", 7, 5, "9919199"},
		{"11111111", 8, 1, "11111111"},
		{"1111111", 7, 1, "1119111"},
		{extremelyLongGiven, 77543, 58343, extremelyLongWant},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := highestValuePalindrome(tt.s, tt.n, tt.k)
			assert.Equal(t, tt.want, got, tt.s)
		})
	}
}

func readExtremelyLongTestData(t *testing.T) (given string, want string) {
	file, err := os.Open("./test_data/highest_value_palindrome.csv")
	assert.NoError(t, err)
	defer func() {
		err := file.Close()
		assert.NoError(t, err)
	}()

	r := csv.NewReader(bufio.NewReader(file))
	rows, err := r.ReadAll()
	assert.NoError(t, err)

	var arr []string
	for _, row := range rows {
		arr = append(arr, row[0])
	}
	assert.Len(t, arr, 2)
	return arr[0], arr[1]
}
