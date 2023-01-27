package p22200

import (
	"fmt"
	"io"
	"math/big"
	"sort"
)

func Solve22238(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	hp := make([]int, n)
	var dir0 direction
	var slope0 *big.Rat
	for i := 0; i < n; i++ {
		var x, y int64
		_, _ = fmt.Fscan(reader, &x, &y, &hp[i])

		if i == 0 {
			slope0, dir0 = getDir(x, y)
		}
	}
	sort.Ints(hp)

	var totDmg int64
	for i := 0; i < m; i++ {
		var x, y int64
		var dmg int64
		_, _ = fmt.Fscan(reader, &x, &y, &dmg)

		slope1, dir1 := getDir(x, y)

		if dir0 != dir1 {
			_, _ = fmt.Fprintln(writer, len(hp))
			continue
		}

		if slope0 != nil && slope1 != nil && slope0.Cmp(slope1) != 0 {
			_, _ = fmt.Fprintln(writer, len(hp))
			continue
		}

		totDmg += dmg
		cutPos := 0
		for cutPos < len(hp) && hp[cutPos] <= int(totDmg) {
			cutPos++
		}

		if cutPos == len(hp) {
			hp = hp[:0]
		} else {
			hp = hp[cutPos:]
		}
		_, _ = fmt.Fprintln(writer, len(hp))
	}
}

func getDir(x, y int64) (slope *big.Rat, dir direction) {
	if x == 0 {
		if y > 0 {
			dir = up
		} else {
			dir = down
		}
		return
	}
	slope = big.NewRat(y, x)
	switch slope.Sign() {
	case 1:
		if x > 0 {
			dir = upperRight
		} else {
			dir = lowerLeft
		}
	case -1:
		if x > 0 {
			dir = lowerRight
		} else {
			dir = upperLeft
		}
	case 0:
		if x > 0 {
			dir = right
		} else {
			dir = left
		}
	}
	return
}

type direction int

const (
	up direction = iota
	upperRight
	right
	lowerRight
	down
	lowerLeft
	left
	upperLeft
)
