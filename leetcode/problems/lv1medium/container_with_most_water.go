package lv1medium

func maxArea(heights []int) int {
	var area int
	for i := 0; i < len(heights)-1; i++ {
		v0 := heights[i]
		for j := i + 1; j < len(heights); j++ {
			width := j - i
			v1 := heights[j]
			if v0 > v1 {
				a := v1 * width
				if a > area {
					area = a
				}
			} else {
				a := v0 * width
				if a > area {
					area = a
				}
			}
		}
	}
	return area
}
