package p10400

import (
	"fmt"
)

func Solve10430(a, b, c int) {
	fmt.Println((a + b) % c)
	fmt.Println((a%c + b%c) % c)
	fmt.Println((a * b) % c)
	fmt.Println((a % c * b % c) % c)
}
