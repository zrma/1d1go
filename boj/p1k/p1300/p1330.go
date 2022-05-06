package p1300

func Solve1330(a, b int) string {
	if a > b {
		return ">"
	}
	if a < b {
		return "<"
	}
	return "=="
}
