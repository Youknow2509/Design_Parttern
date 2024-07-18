
package c

import (
	"v/Interfaces"
	"fmt"
)

/**
 * Concrete implementation of Decorator Plant
 */
type IcePlant struct {
	Plant iface.Plant
	Name string
}

// Constructor
func NewIcePlant(p iface.Plant) *IcePlant {
	return &IcePlant{
		Plant: p,
		Name: "Ice Plant" + p.GetName(),
	}
}

// Deloy method implementation
func (i *IcePlant) Attack() {
	i.Plant.Attack()
	fmt.Println("\t and slow")
}