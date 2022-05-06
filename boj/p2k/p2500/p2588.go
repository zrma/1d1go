package p2500

func Solve2588(a, b int) [4]int {
	third := a * (b % 10)
	fourth := a * (b / 10 % 10)
	fifth := a * (b / 100 % 10)

	return [4]int{
		third,
		fourth,
		fifth,
		third + fourth*10 + fifth*100,
	}
}
