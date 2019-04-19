package tutorial30daysofcode

const (
	weird    = "Weird"
	notWeird = "Not Weird"
)

func cond(n int32) string {
	if n%2 != 0 || (n >= 6 && n <= 20) {
		return weird
	}

	return notWeird
}
