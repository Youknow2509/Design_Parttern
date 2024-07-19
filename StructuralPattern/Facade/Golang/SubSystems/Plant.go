/**
 * Structure Plant
 */
package subsystems

import (
	"fmt"
)

type Plant struct {
}

// Constructor
func NewPlant() *Plant {
	return &Plant{}
}

func (p Plant) Grow() {
	fmt.Println("Plant is growing")
}

func (p Plant) Attack() {
	fmt.Println("Plant is attacking")
}