package strings

func countMap(s string) map[int32]int32 {
	m1 := make(map[int32]int32)
	m2 := make(map[int32]int32)

	for _, b := range s {
		m1[b]++
	}

	for _, v := range m1 {
		m2[v]++
	}

	return m2
}

func valid(s string) bool {
	m := countMap(s)
	if len(m) > 2 {
		return false
	}

	if len(m) == 1 {
		return true
	}

	arr := [2]struct {
		key, value int32
	}{}
	var i int
	for k, v := range m {
		if k == 1 && v == 1 {
			return true
		}

		arr[i] = struct{ key, value int32 }{key: k, value: v}
		i++
	}

	if abs(arr[0].key-arr[1].key) != 1 {
		return false
	}
	return arr[0].value == 1 || arr[1].value == 1
}

func abs(n int32) int32 {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func isValid(s string) string {
	if valid(s) {
		return "YES"
	} else {
		return "NO"
	}
}
