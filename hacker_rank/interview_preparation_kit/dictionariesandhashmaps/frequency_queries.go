package dictionariesandhashmaps

func freqQuery(queries [][]int32) []int32 {
	m1 := make(map[int32]int32) // 횟수 집계 용도
	m2 := make(map[int32]int32) // 횟수 색인 용도

	var result []int32
	for _, q := range queries {
		opCode, opLand := q[0], q[1]
		switch opCode {
		case 1:
			cnt := m1[opLand]
			m1[opLand] = cnt + 1
			m2[cnt]--
			m2[cnt+1]++
		case 2:
			if cnt, ok := m1[opLand]; ok && cnt > 0 {
				m1[opLand]--
				m2[cnt]--
				m2[cnt-1]++
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
