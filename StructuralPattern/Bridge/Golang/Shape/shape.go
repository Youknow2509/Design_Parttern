package shape

import (
	"fmt"
	c "v/color"
)

/**
 * Interface for shape
 */
type Shape interface {
	Draw() 
}

/**
 * Circle shape
 */
type Circle struct {
	color c.Color
}

// Constructor
func NewCircle(color c.Color) *Circle {
	return &Circle{color}
}

// Deloy draw 
func (circle *Circle) Draw() {
	fmt.Println("Circle: ", circle.color.GetColor())
}

/**
 * Square shape
 */
type Square struct {
	color c.Color
}

// Constructor
func NewSquare(color c.Color) *Square {
	return &Square{color}
}

// Deloy draw 
func (square *Square) Draw() {
	fmt.Println("Circle: ", square.color.GetColor())
}

