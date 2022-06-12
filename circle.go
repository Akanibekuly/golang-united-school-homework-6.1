package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (t *Circle) CalcPerimeter() float64 {
	return 2 * math.Pi * t.Radius
}

func (t *Circle) CalcArea() float64 {
	return math.Pi * t.Radius * t.Radius
}
