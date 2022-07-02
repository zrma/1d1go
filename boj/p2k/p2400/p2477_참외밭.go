package p2400

import (
	"fmt"
)

func Solve2477(reader Reader, writer Writer) {
	var ratio int
	_, _ = fmt.Fscan(reader, &ratio)

	const pointCount = 6
	prev := point{}
	points := [pointCount]point{}
	for i := range points {
		var direction, distance int
		_, _ = fmt.Fscan(reader, &direction, &distance)
		switch direction {
		case 1:
			prev.x += distance
			points[i] = prev
		case 2:
			prev.x -= distance
			points[i] = prev
		case 3:
			prev.y -= distance
			points[i] = prev
		default:
			prev.y += distance
			points[i] = prev
		}
	}

	min := points[0]
	max := points[0]
	for _, p := range points {
		if p.x < min.x {
			min.x = p.x
		}
		if p.x > max.x {
			max.x = p.x
		}
		if p.y < min.y {
			min.y = p.y
		}
		if p.y > max.y {
			max.y = p.y
		}
	}

	//   min.x             max.x
	//     |                |
	//   leftTop---------rightTop----- max.y
	//     |                |
	//     |     inside     |
	//     |                |
	//   leftBottom-----rightBottom--- min.y
	//
	var (
		leftTop     point
		rightTop    point
		leftBottom  point
		rightBottom point
		inside      point
	)
	for _, p := range points {
		if p.x != min.x && p.x != max.x && p.y != min.y && p.y != max.y {
			inside = p
		} else if p.x == min.x && p.y == min.y {
			leftBottom = p
		} else if p.x == max.x && p.y == min.y {
			rightBottom = p
		} else if p.x == min.x && p.y == max.y {
			leftTop = p
		} else if p.x == max.x && p.y == max.y {
			rightTop = p
		}
	}
	unknown := point{
		x: leftTop.x ^ rightTop.x ^ leftBottom.x ^ rightBottom.x,
		y: leftTop.y ^ rightTop.y ^ leftBottom.y ^ rightBottom.y,
	}
	outsideArea := (max.x - min.x) * (max.y - min.y)
	insideArea := (inside.x - unknown.x) * (inside.y - unknown.y)
	if insideArea < 0 {
		insideArea = -insideArea
	}
	_, _ = fmt.Fprint(writer, (outsideArea-insideArea)*ratio)
}

type point struct {
	x, y int
}
