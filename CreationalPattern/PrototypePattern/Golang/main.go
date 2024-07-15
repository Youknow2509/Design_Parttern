/**
@author: Ly Tran Vinh
@contact: lytranvinh.work@gmail.com
@content: Prototype Pattern
*/

package main

import (
	"fmt"
	"strconv"
)

/**
 * Interface Plant
 */
type Plant interface {
	ToString()	string
	Clone() 	Plant
	// TODO - Add more methods
}

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
func NewSunFlower(prototype ...interface{}) Plant {
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
func (s *Sunflower) Clone() Plant {
	return &Sunflower{
		Name:   s.Name,
		Price:  s.Price,
		Health: s.Health,
	}
}

func main() {

	// Create Sunflower
	sunFlower := NewSunFlower()
	stringSunFlower := sunFlower.ToString()
	fmt.Println(stringSunFlower)
	fmt.Println("Location memory of Sunflower: ", &sunFlower)

	// Create Clone Sunflower
	cloneSunFlower := sunFlower.Clone()
	// Change properties of Clone Sunflower
	if cloneSunFlower, ok := cloneSunFlower.(*Sunflower); ok {
		cloneSunFlower.Name = "Clone Sunflower"
		cloneSunFlower.Price = 2000
		cloneSunFlower.Health = 300
	} 

	stringCloneSunFlower := cloneSunFlower.ToString()
	fmt.Println(stringCloneSunFlower)
	fmt.Println("Location memory of Clone Sunflower: ", &cloneSunFlower)

}
