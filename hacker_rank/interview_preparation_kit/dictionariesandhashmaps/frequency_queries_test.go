package dictionariesandhashmaps

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFreqQuery(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/frequency-queries/problem")

	for i, tt := range []struct {
		queries [][]int32
		want    []int32
	}{
		{
			[][]int32{
				{1, 5},
				{1, 6},
				{3, 2},
				{1, 10},
				{1, 10},
				{1, 6},
				{2, 5},
				{3, 2},
			},
			[]int32{0, 1},
		},
		{
			[][]int32{
				{3, 4},
				{2, 1003},
				{1, 16},
				{3, 1},
			},
			[]int32{0, 1},
		},
		{
			[][]int32{
				{1, 3},
				{2, 3},
				{3, 2},
				{1, 4},
				{1, 5},
				{1, 5},
				{1, 4},
				{3, 2},
				{2, 4},
				{3, 2},
			},
			[]int32{0, 1, 1},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := freqQuery(tt.queries)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFreqQueryPerformance(t *testing.T) {
	for _, tt := range []struct {
		queriesFile string
		queriesLen  int
		wantFile    string
		wantLen     int
	}{
		{
			queriesFile: "./test_data/frequency_queries_0.csv",
			queriesLen:  1000000,
			wantFile:    "./test_data/frequency_queries_result_0.csv",
			wantLen:     332855,
		},
		{
			queriesFile: "./test_data/frequency_queries_1.csv",
			queriesLen:  100000,
			wantFile:    "./test_data/frequency_queries_result_1.csv",
			wantLen:     33246,
		},
	} {
		t.Run(tt.queriesFile, func(t *testing.T) {
			queries, err := readInputCSV(tt.queriesFile)
			assert.NoError(t, err)
			assert.Len(t, queries, tt.queriesLen)

			want, err := readResultCSV(tt.wantFile)
			assert.NoError(t, err)
			assert.Len(t, want, tt.wantLen)

			assert.Eventually(t, func() bool {
				got := freqQuery(queries)
				return assert.Equal(t, want, got)
			}, time.Second*3, time.Millisecond*100, "시간 초과")
		})
	}
}

func readInputCSV(fileName string) (arr [][]int32, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = file.Close()
	}()

	r := csv.NewReader(bufio.NewReader(file))
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

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

func readResultCSV(fileName string) (arr []int32, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = file.Close()
	}()

	r := csv.NewReader(bufio.NewReader(file))
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		num, err := strconv.ParseInt(strings.TrimSpace(row[0]), 10, 32)
		if err != nil {
			return nil, err
		}
		arr = append(arr, int32(num))
	}
	return arr, nil
}
