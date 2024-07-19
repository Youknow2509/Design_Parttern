/**
 * Structor Zombie
 */
package subsystems

import (
	"fmt"
)

type Zombie struct {
}

// Constructor
func NewZombie() *Zombie {
	return &Zombie{}
}

func (z Zombie) Walk() {
	fmt.Println("Zombie is walking")
}

func (z Zombie) Attack() {
	fmt.Println("Zombie is attacking")
}