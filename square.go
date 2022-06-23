package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	x := s.start.x
	y := s.start.y

	x += int(s.a)
	y += int(s.a)

	return Point{x, y}
}

func (s Square) Area() uint {
	return s.a * s.a
}

func (s Square) Perimeter() uint {
	return s.a * 4
}
