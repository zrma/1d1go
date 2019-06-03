package dp

import (
	"strconv"
)

// 1     : 1                             = 1
// 12    : 1 2 12                        = 15
// 123   : 1 2 12 3 23 123               = 164
// 1234  : 1 2 12 3 23 123 4 34 234 1234 = 1670
//
// 1     : 1
//          |  a     |
// 12    : 1 (2) (12)      ───┐
//            b               │
//                   |       a'       |
// 123   : 1  2   12 | 3  (2)3 (12)3  |  => ((2) + (12)) * 10 + 3*3
//                  b'                   => (0 + 2 + 12) * 10 + 3 + 3 + 3
//                                       => 00 +20 +120       + 3 + 3 + 3
//                                       => 03 +23 +123
// a' = a * 10 + (i + 1) * cur
//
// b' = b + a'
//
// a   0           1            14               149
// a'  1   10+2*2(14)  140+3*3(149)    1490+4*4(1506)
// b   0           1            15               164
// b'  1          15           164              1670

func substrings(n string) int32 {
	const modulo = 1000000007

	length := len(n)
	b, err := strconv.Atoi(string(n[0]))
	if err != nil {
		return 0
	}

	a := b
	var i int
	for i = 1; i < length; i++ {
		cur, err := strconv.Atoi(string(n[i]))
		if err != nil {
			cur = 0
		}
		a = (a*10 + (i+1)*cur) % modulo
		b = (a + b) % modulo
	}

	return int32(b)
}
