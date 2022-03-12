package lv0easy

func romanToInt(s string) int {
	var sum int
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case 'I':
			if i+1 < len(s) {
				if s[i+1] == 'V' {
					sum += 4
					i++
					continue
				}
				if s[i+1] == 'X' {
					sum += 9
					i++
					continue
				}
			}
			sum += 1
		case 'X':
			if i+1 < len(s) {
				if s[i+1] == 'L' {
					sum += 40
					i++
					continue
				}
				if s[i+1] == 'C' {
					sum += 90
					i++
					continue
				}
			}
			sum += 10
		case 'C':
			if i+1 < len(s) {
				if s[i+1] == 'D' {
					sum += 400
					i++
					continue
				}
				if s[i+1] == 'M' {
					sum += 900
					i++
					continue
				}
			}
			sum += 100
		case 'V':
			sum += 5
		case 'L':
			sum += 50
		case 'D':
			sum += 500
		case 'M':
			sum += 1000
		}
	}
	return sum
}
