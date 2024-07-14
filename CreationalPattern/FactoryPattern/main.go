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

// Constructor for SunFlower 
func NewSunFlower(params...interface{}) *SunFlower {
	p := &SunFlower {
		Name: "SunFlower",
		Price: 50,
	}

	len := len(params)

	if (len > 0) {
		p.Name = params[0].(string)
	}

	if (len > 1) {
		p.Price = params[1].(int)
	}

	return p
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

// Constructor for PeaShooter
func NewPeaShooter(params...interface{}) *PeaShooter {
	p := &PeaShooter {
		Name: "PeaShooter",
		Price: 100,
	}

	len := len(params)

	if (len > 0) {
		p.Name = params[0].(string)
	}

	if (len > 1) {
		p.Price = params[1].(int)
	}

	return p
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

// Constructor for WallNut
func NewWallNut(params...interface{}) *WallNut {
	p := &WallNut {
		Name: "WallNut",
		Price: 50,
	}

	len := len(params)

	if (len > 0) {
		p.Name = params[0].(string)
	}

	if (len > 1) {
		p.Price = params[1].(int)
	}

	return p
}

func (s *WallNut) PlantName() string {
	return s.Name
}

func (s *WallNut) PlantPrice() int {
	return s.Price
}

func main() {

	// Create SunFlower
	sunFlower := NewSunFlower("SunFlower", 550)
	fmt.Println("SunFlower: ", sunFlower.PlantName(), " - ", sunFlower.PlantPrice())

	// Create PeaShooter
	peaShooter := NewPeaShooter("PeaShooter", 100)
	fmt.Println("PeaShooter: ", peaShooter.PlantName(), " - ", peaShooter.PlantPrice())

	// Create WallNut
	wallNut := NewWallNut("WallNut", 50)
	fmt.Println("WallNut: ", wallNut.PlantName(), " - ", wallNut.PlantPrice())

	sunFlower2 := NewSunFlower()
	fmt.Println("SunFlower: ", sunFlower2.PlantName(), " - ", sunFlower2.PlantPrice())
	
}
