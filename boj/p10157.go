package boj

type direction int

const (
	up direction = 0 + iota
	right
	down
	left
	dirMax
)

func turn(d direction) direction {
	return (d + 1) % dirMax
}

func Solve10157(col, row, k int) (int, int, bool) {
	if row*col < k {
		return 0, 0, false
	}

	x, y := 1, 1
	k0 := k - 1
	col0, row0 := col-1, row
	dir := up
	for {
		switch dir {
		case up:
			if k0 < row0 {
				return x, y + k0, true
			}
			x += 1
			y += row0 - 1
			k0 -= row0
			row0--
			dir = turn(dir)
		case right:
			if k0 < col0 {
				return x + k0, y, true
			}
			x += col0 - 1
			y -= 1
			k0 -= col0
			col0--
			dir = turn(dir)
		case down:
			if k0 < row0 {
				return x, y - k0, true
			}
			x -= 1
			y -= row0 - 1
			k0 -= row0
			row0--
			dir = turn(dir)
		case left:
			if k0 < col0 {
				return x - k0, y, true
			}
			x -= col0 - 1
			y += 1
			k0 -= col0
			col0--
			dir = turn(dir)
		}
	}
}
