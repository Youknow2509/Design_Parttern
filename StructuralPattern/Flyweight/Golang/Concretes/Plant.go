package Concretes

import (
	"fmt"
	"v/Interfaces"
)

/**
 * Struct Plant implementation Game Object
 */
type Plant struct {
    Type string // Intrinsic state
    Image string
}

// Construct a plant
func NewPlant(t string, image string) (iface.GameObject) {
	return &Plant{Type: t, Image: image}
}

// Override display
func (p *Plant) Display(x int, y int) {
	// Display the plant
	fmt.Println("Displaying Plant: ", p.Type, " at (", x, ", ", y, ")")
}
