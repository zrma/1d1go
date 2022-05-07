package lv0easy

func isPalindrome(n int) bool {
	if n < 0 {
		return false
	}
	if n == 0 {
		return true
	}

	var rev int
	x := n
	for x > 0 {
		y := x % 10
		rev = rev*10 - y
		x = x / 10
	}

	return n+rev == 0
}
