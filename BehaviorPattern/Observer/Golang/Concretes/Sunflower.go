package concretes

import (
	"v/Interfaces"
	"fmt"
)

type Sunflower struct {
	PlantAction string
}

// Constructor
func NewSunflower() interfaces.Observer {
	return &Sunflower{}
}

// Overise update
func (p *Sunflower) Update(action string) {
	p.PlantAction = action
	fmt.Println("Peashooter: Zombie action detected: " + action)
}