package p2500

func Solve2588(fmt InOut) {
	var a, b int
	_, _ = fmt.Scan(&a, &b)

	third := a * (b % 10)
	fourth := a * (b / 10 % 10)
	fifth := a * (b / 100 % 10)

	_, _ = fmt.Println(third)
	_, _ = fmt.Println(fourth)
	_, _ = fmt.Println(fifth)
	_, _ = fmt.Println(third + fourth*10 + fifth*100)
}
