package golang_united_school_homework

import (
	"fmt"
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
		return fmt.Errorf("goes out of the shapesCapacity range")
	}

	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, fmt.Errorf("index doesn't exist or index went out of the range")
	}
	return b.shapes[i], nil

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, fmt.Errorf("index doesn't exist or index went out of the range")
	}

	shape := b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, fmt.Errorf("index doesn't exist or index went out of the range")
	}
	removed := b.shapes[i]
	b.shapes[i] = shape
	return removed, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	count := 0
	var i interface{} = ""
	result := []Shape{}
	for _, r := range b.shapes {
		i = r
		_, ok := i.(*Circle)
		fmt.Println(r)
		if ok {
			count++

		} else {
			result = append(result, r)

		}
	}

	if count == 0 {
		return fmt.Errorf("circles are not exist in the list")
	}
	b.shapes = result
	return nil

}

func (b *box) SumPerimeter() float64 {
	result := float64(0)
	for _, num := range b.shapes {
		result += float64(num.CalcPerimeter())
	}
	return result
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	result := float64(0)
	for _, num := range b.shapes {
		result += float64(num.CalcArea())
	}
	return result
}
