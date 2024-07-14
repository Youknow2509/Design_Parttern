/*
*

	@author: Ly Tran Vinh
	@contact: lytranvinh.work@gmail.com
	@content: Observer Pattern
*/
package main

import (
	"fmt"
)

type TypePlant int

/**
 * Enum Type PlantType
 */

const (
	SUNFLOWER TypePlant = iota
	PEASHOOTER
	WALLNUT
	// TODO add more plant type ...
)

// String method for TypePlant type
func (d TypePlant) String() string {
	return [...]string{
		"SunFlower",
		"PeaShooter",
		"WallNut",
		// TODO add more plant type ...
	}[d]
}

/**
 * Plant interface
 */
type Plant interface {
	PlantName() string
	PlantPrice() int
	// TODO add more method ...
}

/**
 * SunFlower struct, implement Plant interface
 */
type SunFlower struct {
	Name  string
	Price int
	// TODO add more attribute ...
}

func (s *SunFlower) PlantName() string {
	return s.Name
}

func (s *SunFlower) PlantPrice() int {
	return s.Price
}

/**
 * PEASHOOTER struct, implement Plant interface
 */
type PeaShooter struct {
	Name  string
	Price int
	// TODO add more attribute ...
}

func (s *PeaShooter) PlantName() string {
	return s.Name
}

func (s *PeaShooter) PlantPrice() int {
	return s.Price
}

/**
 * WallNut struct, implement Plant interface
 */
type WallNut struct {
	Name  string
	Price int
	// TODO add more attribute ...
}

func (s *WallNut) PlantName() string {
	return s.Name
}

func (s *WallNut) PlantPrice() int {
	return s.Price
}

func main() {

	fmt.Println(SUNFLOWER)
	fmt.Println(PEASHOOTER)

}
