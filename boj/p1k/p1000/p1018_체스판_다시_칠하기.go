package p1000

import (
	"fmt"
	"io"
	"math"
)

func Solve1018(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	board := make([]string, n)
	for i := range board {
		var s string
		_, _ = fmt.Fscan(reader, &s)

		board[i] = s[:m]
	}

	res := calcMinPaintingCount(board, m, n)
	_, _ = fmt.Fprint(writer, res)
}

const maxLen = 8

func calcMinPaintingCount(board []string, width, height int) int {
	minPaintingCount := math.MaxInt32
	for y := 0; y <= height-maxLen; y++ {
		for x := 0; x <= width-maxLen; x++ {
			if paintingCount := calcPaintingCount(board, x, y); paintingCount < minPaintingCount {
				minPaintingCount = paintingCount
			}
		}
	}
	return minPaintingCount
}

func calcPaintingCount(board []string, offsetX, offsetY int) int {
	//goland:noinspection SpellCheckingInspection
	const (
		type0 = "WBWBWBWB"
		type1 = "BWBWBWBW"
	)
	var (
		diffCount0 = 0
		diffCount1 = 0
	)
	for y := offsetY; y < offsetY+maxLen; y++ {
		targetLine := board[y][offsetX : offsetX+maxLen]
		for i := range targetLine {
			c := targetLine[i]
			if y%2 == 0 {
				if c != type0[i] {
					diffCount0++
				}
				if c != type1[i] {
					diffCount1++
				}
			} else {
				if c != type1[i] {
					diffCount0++
				}
				if c != type0[i] {
					diffCount1++
				}
			}
		}
	}
	if diffCount0 < diffCount1 {
		return diffCount0
	} else {
		return diffCount1
	}
}
