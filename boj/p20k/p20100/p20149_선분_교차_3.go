package p20100

import (
	"fmt"
	"io"
	"math"

	"1d1go/boj/p17k/p17300"
)

func Solve20149(reader io.Reader, writer io.Writer) {
	const pointCount = 4
	var points [pointCount]p17300.Point
	for i := range points {
		_, _ = fmt.Fscan(reader, &points[i].X, &points[i].Y)
	}

	equalLine, ok := p17300.IsIntersect(points[0], points[1], points[2], points[3])
	if ok {
		_, _ = fmt.Fprintln(writer, 1)

		intersectPos(points[0], points[1], points[2], points[3], equalLine, writer)
	} else {
		_, _ = fmt.Fprintln(writer, 0)
	}
}

func intersectPos(p0, p1, p2, p3 p17300.Point, equalLine bool, writer io.Writer) {
	x1, y1, x2, y2 := float64(p0.X), float64(p0.Y), float64(p1.X), float64(p1.Y)
	x3, y3, x4, y4 := float64(p2.X), float64(p2.Y), float64(p3.X), float64(p3.Y)

	if equalLine {
		const epsilon = 1e-9

		// // IsIntersect returns whether two lines intersect
		//
		// | Step 1. 왼쪽에서 오른쪽 아래로 내려가며 곱해서 더하기
		// |
		// |                    x1     y1
		// |                        ↘
		// |                    x2     y2  +x1y2
		// |                        ↘
		// |                    x3     y3  +x2y3
		// |                        ↘
		// |                    x1     y1  +x3y1
		// |                ---------------------
		// |                               +x1y2 +x2y3 +x3y1
		// |
		// |  Step 2. 오른쪽에서 왼쪽 아래로 내려가며 곱해서 빼기
		// |
		// |                    x1     y1
		// |                        ↙
		// |             -y1x2  x2     y2
		// |                        ↙
		// |             -y2x3  x3     y3
		// |                        ↙
		// |             -y3x1  x1     y1
		// |                 ---------------------
		// |  -y1x2 -y2x3 -y3x1
		// |
		// |  (x1y2 + x2y3 + x3y1 - y1x2 - y2x3 - y3x1)
		//
		// 두 선분 모두, 다른 선분의 두 점에 대해 두 CCW 값의 부호가 서로 반대라면 두 선분은 교차한다.
		//
		// CCW 값이 0(삼각형의 면적이 0)이면 세 점이 일직선 위에 있다는 것을 알 수 있는데
		// *모든 외적 값이 모두 0이면 두 선분을 확장한 직선이 **겹치는 상황**이고*
		// 두 삼각형 중 한 삼각형만 외적 값이 0이면 한 선분의 끝점에 다른 선분을 확장한 직선이 걸쳐지는 상황이다.
		//
		// 두 직선이 겹칠 경우에는 - (A)
		//
		// - 각 선분의 길이의 합 > 두 선분 전체의 길이 (겹치지 않음)
		// - 각 선분의 길이의 합 == 두 선분 전체의 길이 (끝점만 겹침, 단 하나의 교점)
		// - 각 선분의 길이의 합 < 두 선분 전체의 길이 (두 직선이 여러 점에서 겹침, 부정)
		//
		// 이렇게 선분의 길이를 통해 단 한점에서 교차하는지를 알 수 있다.
		//
		// 두 선분이 겹치고 단 한점에서 교차할 경우는 끝점 비교만으로 교차점을 찾을 수 있다.
		if math.Abs(p0.Distance(p1)+p2.Distance(p3)-p17300.MaxDistance(p0, p1, p2, p3)) <= epsilon {
			if y1 == y3 || y1 == y4 {
				_, _ = fmt.Fprint(writer, x1, y1)
			} else if y2 == y3 || y2 == y4 {
				_, _ = fmt.Fprint(writer, x2, y2)
			}
		}
		return
	}

	// 두 선분이 겹치지 않고 교차할 경우는 1차 함수의 연립방정식을 이용해 아래와 같이 교차점을 찾는다.
	//
	// (y1 = ax1 + b)
	// (y2 = ax2 + b)
	//
	// (y3 = cx3 + d)
	// (y4 = cx4 + d)
	//
	// (y = ax + b) - (B)
	//
	// (y1 = ax1 + b)
	// (y2 = ax2 + b)
	// (y1 - y2 = ax1 - ax2)
	// (y1 - y2 = a(x1 - x2))
	// (x(y1 - y2) = ax(x1 - x2))
	//
	// (y(x1 - x2) = ax(x1 - x2) + b(x1 - x2)) - (B)의 양 변에 (x1 - x2) 곱함
	// (y(x1 - x2) = x(y1 - y2) + b(x1 - x2)) - (C)
	//
	// (y1(x1 - x2) = ax1(x1 - x2) + b(x1 - x2))
	// (y2(x1 - x2) = ax2(x1 - x2) + b(x1 - x2))
	// ((y1 + y2)(x1 - x2) = a(x1 + x2)(x1 - x2) + 2b(x1 - x2))
	// ((x1 - x2)(y1 + y2) = a(x1 - x2)(x1 + x2) + 2b(x1 - x2))
	// ((x1 - x2)(y1 + y2) = (y1 - y2)(x1 + x2) + 2b(x1 - x2))
	// ((x1 - x2)(y1 + y2) = (x1 + x2)(y1 - y2) + 2b(x1 - x2))
	// ((x1 - x2)(y1 + y2) - (x1 + x2)(y1 - y2) = 2b(x1 - x2))
	// (x1y1 + x1y2 - x2y1 - x2y2 - x1y1 + x1y2 - x2y1 + x2y2 = 2b(x1 - x2))
	// (x1y2 - x2y1 + x1y2 - x2y1 = 2b(x1 - x2))
	// (2x1y2 - 2x2y1 = 2b(x1 - x2))\
	// (x1y2 - x2y1 = b(x1 - x2)) - (D)
	//
	// (y(x1 - x2) = x(y1 - y2) + x1y2 - x2y1) - (C)에 (D) 대입
	//
	// 이 식에서 역으로 직선 위의 (x, y), (x1, y1), (x2, y2) 세 점에 대해 CCW가 0임을 유도할 수 있다.
	//
	// (y(x1 - x2) = xy1 - xy2 + x1y2 - x2y1)
	// (y(x1 - x2) = x1y2 + xy1 - y1x2 - y2x)
	// (yx1 - x2y = x1y2 + xy1 - y1x2 - y2x)
	// (x1y2 + x2y + xy1 - y1x2 - y2x - yx1 = 0) - [Shoelace formula](https://en.wikipedia.org/wiki/Shoelace_formula)
	//
	// 마찬가지 방식으로 x3, y3에 대해서도 아래와 같이 식을 유도할 수 있다.
	//
	// (y(x3 - x4) = x(y3 - y4) + x3y4 - x4y3)
	//
	// (y(x1 - x2)(x3 - x4) = x(y1 - y2)(x3 - x4) + x1y2(x3 - x4) - x2y1(x3 - x4))
	// (y(x1 - x2)(x3 - x4) = x(x1 - x2)(y3 - y4) + x3y4(x1 - x2) - x4y3(x1 - x2))
	//
	// (0 = x(y1 - y2)(x3 - x4) + x1y2(x3 - x4) - x2y1(x3 - x4) - x(x1 - x2)(y3 - y4) - x3y4(x1 - x2) + x4y3(x1 - x2))
	// (0 = x(y1 - y2)(x3 - x4) - x(x1 - x2)(y3 - y4) + x1y2(x3 - x4) - x2y1(x3 - x4) - x3y4(x1 - x2) + x4y3(x1 - x2))
	// (x(x1 - x2)(y3 - y4) - x(y1 - y2)(x3 - x4) = x1y2(x3 - x4) - x2y1(x3 - x4) - x3y4(x1 - x2) + x4y3(x1 - x2))
	// (x((x1 - x2)(y3 - y4) - (y1 - y2)(x3 - x4)) = (x1y2 - x2y1)(x3 - x4) - (x3y4 - x4y3)(x1 - x2))
	//
	// (x1 - x2)(y3 - y4) - (y1 - y2)(x3 - x4) == 0 인 경우는
	// 두 직선의 기울기가 같거나(CCW 체크 실패) 두 직선이 겹치는 경우로 위에서 별도로 처리(A)하거나,
	// 선분의 길이가 0보다 크다는 문제의 가정에 위배된다.
	//
	// (x = ((x1y2 - x2y1)(x3 - x4) - (x3y4 - x4y3)(x1 - x2)) / ((x1 - x2)(y3 - y4) - (y1 - y2)(x3 - x4)))
	//
	// 마찬가지 방식으로 y를 구하면 다음과 같다.
	//
	// (y = ((x1y2 - x2y1)(y3 - y4) - (x3y4 - x4y3)(y1 - y2)) / ((x1 - x2)(y3 - y4) - (y1 - y2)(x3 - x4)))
	x := ((x1*y2-x2*y1)*(x3-x4) - (x3*y4-x4*y3)*(x1-x2)) / ((x1-x2)*(y3-y4) - (y1-y2)*(x3-x4))
	y := ((x1*y2-x2*y1)*(y3-y4) - (x3*y4-x4*y3)*(y1-y2)) / ((x1-x2)*(y3-y4) - (y1-y2)*(x3-x4))
	_, _ = fmt.Fprint(writer, x, y)
}
