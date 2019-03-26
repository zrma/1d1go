package dictionaries_and_hashmaps

func freqQuery(queries [][]int32) []int32 {
	m1 := make(map[int32]int32)
	m2 := make(map[int32]int32)

	var result []int32
	for _, q := range queries {
		opCode, opLand := q[0], q[1]
		switch opCode {
		case 1:
			old := m1[opLand]
			m1[opLand] = old + 1
			m2[old]--
			m2[old+1]++
		case 2:
			if old, ok := m1[opLand]; ok && old > 0 {
				m1[opLand]--
				m2[old]--
				m2[old-1]++
			}
		case 3:
			if m2[opLand] > 0 {
				result = append(result, 1)
			} else {
				result = append(result, 0)
			}
		}
	}

	return result
}
