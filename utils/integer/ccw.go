package integer

// CCW returns the direction of the cross product of the vectors (x1, y1) -> (x2, y2) and (x2, y2) -> (x3, y3)
// 1: counter-clockwise
// -1: clockwise
// 0: collinear
//
// Shoelace formula(신발끈 공식, 가우스의 면적 공식, 측량사의 공식)
// https://en.wikipedia.org/wiki/Shoelace_formula
//
// 외적의 수직벡터 방향에 따라 반시계방향의 삼각형의 면적은 양수, 시계방향의 경우는 음수
//
// 오른손 법칙을 기억하자
func CCW(x1, y1, x2, y2, x3, y3 int) int {
	result := x1*y2 + x2*y3 + x3*y1 - y1*x2 - y2*x3 - y3*x1
	if result > 0 {
		return 1
	} else if result < 0 {
		return -1
	}
	return 0
}
