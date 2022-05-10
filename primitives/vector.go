package primitives

import "math"
import "math/rand"

type Vector struct {
	X, Y float64
}

func (v *Vector) Add(toAdd *Vector) *Vector {
	v.X += toAdd.X
	v.Y += toAdd.Y
	return v
}

func (v *Vector) MultiplyBy(multiplier float64) *Vector {
	v.X *= multiplier
	v.Y *= multiplier
	return v
}

func (v *Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector) Normalize() *Vector {
	length := v.Length()
	if length == 0 {
		return v
	}
	v.X /= length
	v.Y /= length
	return v
}

func (v *Vector) Negate() *Vector {
	v.X = -v.X
	v.Y = -v.Y
	return v
}

func NegativeVector(v *Vector) *Vector {
	return &Vector{-v.X, -v.Y}
}

func AddVectors(v1, v2 *Vector) *Vector {
	return &Vector{v1.X + v2.X, v1.Y + v2.Y}
}

func MultiplyVectorBy(v *Vector, multiplier float64) *Vector {
	return &Vector{v.X * multiplier, v.Y * multiplier}
}

func NewVector(x, y float64) *Vector {
	return &Vector{x, y}
}

func RandomNormalizedVector() *Vector {
	vector := Vector{
		X: rand.Float64()*2 - 1,
		Y: rand.Float64()*2 - 1,
	}
	return vector.Normalize()
}
