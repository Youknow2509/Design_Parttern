package Concretes

import (
	"fmt"
	"v/Interfaces"
)

/**
 * Struct Zombie implementation Game Object
 */
type Zombie struct {
    Type string // Intrinsic state
    Image string
}

// Construct a Zombie
func NewZombie(t string, image string) (iface.GameObject) {
	return &Zombie{Type: t, Image: image}
}

// Override display
func (z *Zombie) Display(x int, y int) {
	// Display the Zombie
	fmt.Println("Displaying Zombie: ", z.Type, " at (", x, ", ", y, ")")
}
