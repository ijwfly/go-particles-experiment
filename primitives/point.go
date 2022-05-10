package primitives

import "math"

type Point struct {
	X, Y float64
}

func (p *Point) Update(x, y float64) {
	p.X = x
	p.Y = y
}

func (p *Point) MoveByVector(v *Vector) {
	p.X += v.X
	p.Y += v.Y
}

func (p *Point) DistanceTo(toPoint *Point) float64 {
	return math.Sqrt(math.Pow(toPoint.X-p.X, 2) + math.Pow(toPoint.Y-p.Y, 2))
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}
