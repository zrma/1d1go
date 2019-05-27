package strings

func morganAndString(a string, b string) string {
	var i, j int
	var output []byte
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			output = append(output, a[i])
			i++
			continue
		}

		if a[i] > b[j] {
			output = append(output, b[j])
			j++
			continue
		}

		// 동률일 경우 다음 자리 비교
		if compare(a, b, i+1, j+1) {
			output, i = extend(a, output, i)
			continue
		}

		output, j = extend(b, output, j)
	}

	if i < len(a) {
		output = append(output, a[i:]...)
	}

	if j < len(b) {
		output = append(output, b[j:]...)
	}

	return string(output)
}

func extend(in string, out []byte, i int) ([]byte, int) {
	out = append(out, in[i])
	i++

	// 반복 문자열은 일괄 붙이기
	for i < len(in) && in[i] == in[i-1] {
		out = append(out, in[i])
		i++
	}

	return out, i
}

func compare(a, b string, i, j int) bool {
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			return true
		}

		if a[i] > b[j] {
			return false
		}

		i++
		j++
	}

	return i != len(a)
}
