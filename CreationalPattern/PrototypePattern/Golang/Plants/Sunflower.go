
package Plants

import (
	"strconv"
	Interface "v/interface"
)

/**
 * Struct Sunflower implement Plant interface
 */
type Sunflower struct {
	Name   string
	Price  int
	Health int
	// TODO - Add more properties
}

// Constructor of Sunflower
func NewSunFlower(prototype ...interface{}) Interface.Plant {
	p := &Sunflower{
		Name:   "Sunflower",
		Price:  1000,
		Health: 200,
	}

	if len(prototype) > 0 {
		if prototype[0] != nil {
			p.Name = prototype[0].(string)
		}
		if prototype[1] != nil {
			p.Price = prototype[1].(int)
		}
		if prototype[2] != nil {
			p.Health = prototype[2].(int)
		}
	}

	return p
}

// Deloyment ToString method of Sunflower
func (s *Sunflower) ToString() string {
	return "Name: " + s.Name + ", Price: " + strconv.Itoa(s.Price) + ", Health: " + strconv.Itoa(s.Health)
}

// Deloyment Clone method of Sunflower
func (s *Sunflower) Clone() Interface.Plant {
	return &Sunflower{
		Name:   s.Name,
		Price:  s.Price,
		Health: s.Health,
	}
}