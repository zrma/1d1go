package strings

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPalindromeIndex(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/palindrome-index/problem")

	//noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want int32
	}{
		{"aaab", 3},
		{"baa", 0},
		{"aaa", -1},
		{"a", -1},
		{"abcadcdddaba", -1},
		{"abcd", -1},
		{"quyjjdcgsvvsgcdjjyq", 1},
		{"hgygsvlfwcwnswtuhmyaljkqlqjjqlqkjlaymhutwsnwcflvsgygh", 8},
		{"fgnfnidynhxebxxxfmxixhsruldhsaobhlcggchboashdlurshxixmfxxxbexhnydinfngf", 33},
		{"bsyhvwfuesumsehmytqioswvpcbxyolapfywdxeacyuruybhbwxjmrrmjxwbhbyuruycaexdwyfpaloyxbcpwsoiqtymhesmuseufwvhysb", 23},
		{"fvyqxqxynewuebtcuqdwyetyqqisappmunmnldmkttkmdlnmnumppasiqyteywdquctbeuwenyxqxqyvf", 25},
		{"mmbiefhflbeckaecprwfgmqlydfroxrblulpasumubqhhbvlqpixvvxipqlvbhqbumusaplulbrxorfdylqmgfwrpceakceblfhfeibmm", 44},
		{"tpqknkmbgasitnwqrqasvolmevkasccsakvemlosaqrqwntisagbmknkqpt", 20},
		{"lhrxvssvxrhl", -1},
		{"prcoitfiptvcxrvoalqmfpnqyhrubxspplrftomfehbbhefmotfrlppsxburhyqnpfmqlaorxcvtpiftiocrp", 14},
		{"kjowoemiduaaxasnqghxbxkiccikxbxhgqnsaxaaudimeowojk", -1},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := palindromeIndex(tt.give)
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestEqualPrefix(t *testing.T) {
	//noinspection SpellCheckingInspection
	for i, tt := range []struct {
		description string
		s1, s2      string
		want        bool
	}{
		{
			description: "single-character string, true if equal",
			s1:          "a",
			s2:          "a",
			want:        true,
		},
		{
			description: "single-character string, false if not equal",
			s1:          "a",
			s2:          "b",
			want:        false,
		},
		{
			description: "over two characters string, true if the first two characters are the same",
			s1:          "abcd",
			s2:          "abab",
			want:        true,
		},
		{
			description: "over two characters string, false if the first two characters are not the same",
			s1:          "abc",
			s2:          "bbc",
			want:        false,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := equalPrefix(tt.s1, tt.s2)
			assert.Equal(t, tt.want, got, tt.description)
		})
	}
}

func TestPalindromeIndexPerformance(t *testing.T) {
	file, err := os.Open("./test_data/palindrome_index.csv")
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

	want := []int32{
		16722,
		76661,
		10035,
		57089,
		46674,
	}

	assert.Len(t, arr, 5)
	for i, s := range arr {
		assert.Eventually(t, func() bool {
			got := palindromeIndex(s)
			return assert.EqualValues(t, want[i], got)
		}, time.Second, time.Millisecond*100, "시간 초과")
	}
}
