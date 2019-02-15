package tutorial_30_days_of_code

const (
	Weird    = "Weird"
	NotWeird = "Not Weird"
)

func cond(n int32) string {
	if n%2 != 0 || (n >= 6 && n <= 20) {
		return Weird
	}

	return NotWeird
}
