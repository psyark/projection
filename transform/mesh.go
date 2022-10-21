package transform

type Point struct {
	X, Y float64
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

func (p Point) Dot(q Point) float64 {
	return p.X*q.Y - p.Y*q.X
}
