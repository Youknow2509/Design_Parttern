package concretes

import (
	"v/Interfaces"
	"fmt"
)

type Peashooter struct {
	PlantAction string
}

// Constructor
func NewPeashooter() interfaces.Observer {
	return &Peashooter{}
}

// Overise update
func (p *Peashooter) Update(action string) {
	p.PlantAction = action
	fmt.Println("Peashooter: Zombie action detected: " + action)
}