package dp

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBricksGame(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/play-game/problem")

	for _, tt := range []struct {
		description string
		given       []int32
		want        int64
	}{
		{
			description: "1~3",
			given:       []int32{1, 2, 3},
			want:        6,
		},
		{
			description: "1~4",
			given:       []int32{1, 2, 3, 4},
			want:        6,
		},
		{
			description: "1~5",
			given:       []int32{1, 2, 3, 4, 5},
			want:        6,
		},
		{
			description: "big number front, zero back",
			given:       []int32{999, 1, 1, 1, 0},
			want:        1001,
		},
		{
			description: "zero front, big number end",
			given:       []int32{0, 1, 1, 1, 999},
			want:        999,
		},
		{
			description: "multiple duplicate values",
			given:       []int32{0, 1, 1, 1, 999, 999},
			want:        1001,
		},
	} {
		t.Run(tt.description+"/loop", func(t *testing.T) {
			got := bricksGame(tt.given)
			assert.Equal(t, tt.want, got)
		})

		t.Run(tt.description+"/recur", func(t *testing.T) {
			got := bricksGameRecur(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBricksGame_Reversal(t *testing.T) {
	for _, tt := range []struct {
		description string
		given       []int32
		want        int64
	}{
		{
			description: "1~3",
			given:       []int32{1, 2, 3},
			want:        0,
		},
		{
			description: "1~4",
			given:       []int32{1, 2, 3, 4},
			want:        4,
		},
		{
			description: "1~5",
			given:       []int32{1, 2, 3, 4, 5},
			want:        9,
		},
		{
			description: "big number front, zero back",
			given:       []int32{999, 1, 1, 1, 0},
			want:        1,
		},
		{
			description: "zero front, big number end",
			given:       []int32{0, 1, 1, 1, 999},
			want:        3,
		},
		{
			description: "multiple duplicate values",
			given:       []int32{0, 1, 1, 1, 999, 999},
			want:        1000,
		},
	} {
		t.Run(tt.description, func(t *testing.T) {
			got := bricksGameReversal(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}

func readCSV(tb testing.TB, fileName string) ([][]int32, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := file.Close()
		assert.NoError(tb, err)
	}()

	r := csv.NewReader(bufio.NewReader(file))
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	var arr [][]int32
	for _, row := range rows {
		var cols []int32
		for _, col := range row {
			num, err := strconv.ParseInt(strings.TrimSpace(col), 10, 32)
			if err != nil {
				return nil, err

			}
			cols = append(cols, int32(num))
		}
		arr = append(arr, cols)
	}
	return arr, nil
}

type fixture struct {
	given [][]int32
	want  []struct {
		len int
		val int64
	}
}

func newFixture(tb testing.TB) fixture {
	arr, err := readCSV(tb, "./test_data/bricks_game_0.csv")
	assert.NoError(tb, err)

	arr2, err := readCSV(tb, "./test_data/bricks_game_1.csv")
	assert.NoError(tb, err)

	arr = append(arr, arr2...)
	assert.Len(tb, arr, 8)

	return fixture{
		given: arr,
		want: []struct {
			len int
			val int64
		}{
			{1000, 249147},
			{1000, 251633},
			{1000, 249302},
			{1000, 247737},
			{1000, 253105},
			{100000, 249791261588},
			{100000, 249894676936},
			{100000, 250224672758},
		},
	}
}

func TestBricksGamePerformance(t *testing.T) {
	assert.Eventually(t, func() bool {
		f := newFixture(t)

		for i, arr := range f.given {
			assert.Len(t, arr, f.want[i].len)

			got := bricksGame(arr)
			assert.EqualValues(t, f.want[i].val, got)
		}

		return true
	}, time.Second, time.Millisecond*100, "시간 초과")
}

//goos: windows
//goarch: amd64
//pkg: 1d1go/hacker_rank/algorithms/dp
//cpu: AMD Ryzen 7 3700X 8-Core Processor
//BenchmarkBrickGame
//BenchmarkBrickGame/Loop
//BenchmarkBrickGame/Loop-16                  4346            265899 ns/op
//BenchmarkBrickGame/Recur
//BenchmarkBrickGame/Recur-16                   38          27910642 ns/op
//PASS
func BenchmarkBrickGame(b *testing.B) {
	b.Run("Loop", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			b.StopTimer()
			f := newFixture(b)
			b.StartTimer()

			for pb.Next() {
				for _, arr := range f.given {
					bricksGame(arr)
				}
			}
		})
	})

	b.Run("Recur", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			b.StopTimer()
			f := newFixture(b)
			b.StartTimer()

			for pb.Next() {
				for _, arr := range f.given {
					bricksGameRecur(arr)
				}
			}
		})
	})
}
