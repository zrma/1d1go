package integer

func CCW(x1, y1, x2, y2, x3, y3 int) int {
	result := x1*y2 + x2*y3 + x3*y1 - y1*x2 - y2*x3 - y3*x1
	if result > 0 {
		return 1
	} else if result < 0 {
		return -1
	}
	return 0
}
