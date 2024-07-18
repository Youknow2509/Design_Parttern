/**
 * Struct Peashooter implementation for Plant interface
 */

package c

import (
	"fmt"
)

type Peashooter struct {
	Name string
}

// Constructor
func NewPeashooter() *Peashooter {
	return &Peashooter{
		Name: "Peashooter",
	}
}

// Deloy method implementation
func (p *Peashooter) Attack() {
	fmt.Printf("%v is attacking\n", p.Name)
}

// Deloy method implementation
func (p *Peashooter) GetName() string {
	return p.Name
}
