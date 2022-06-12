package golang_united_school_homework

import (
	"fmt"
	"reflect"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) == b.shapesCapacity {
		return fmt.Errorf("box out of capacity")
	}

	b.shapes = append(b.shapes, shape)

	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, fmt.Errorf("index out of range")
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, fmt.Errorf("index out of range")
	}

	sh := b.shapes[i]

	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)

	return sh, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, fmt.Errorf("index out of range")
	}

	sh := b.shapes[i]
	b.shapes[i] = shape

	return sh, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for _, sh := range b.shapes {
		sum += sh.CalcPerimeter()
	}

	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for _, sh := range b.shapes {
		sum += sh.CalcArea()
	}

	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	sh := make([]Shape, 0, len(b.shapes))
	for i := range b.shapes {
		if !(reflect.TypeOf(b.shapes[i]) == reflect.TypeOf(&Circle{})) {
			sh = append(sh, b.shapes[i])
		}
	}

	if len(sh) == len(b.shapes) {
		return fmt.Errorf("circles are not exist in the list")
	}

	b.shapes = sh

	return nil
}
