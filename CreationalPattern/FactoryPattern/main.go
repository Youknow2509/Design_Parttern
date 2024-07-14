/**
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

/** 
 * Factory method
 */
func CreatePlant(plantType TypePlant) (Plant, error) {
	switch plantType {
	case SUNFLOWER:
		return NewSunFlower(), nil
	case PEASHOOTER:
		return NewPeaShooter(), nil
	case WALLNUT:
		return NewWallNut(), nil
	// TODO add more plant type ...
	default:
		return nil, fmt.Errorf("Invalid plant type")
	}
}

func main() {

	plantType := SUNFLOWER
	plant, err := CreatePlant(plantType)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Plant Name: ", plant.PlantName())
	fmt.Println("Plant Price: ", plant.PlantPrice())

	plantType = PEASHOOTER
	plant, err = CreatePlant(plantType)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Plant Name: ", plant.PlantName())
	fmt.Println("Plant Price: ", plant.PlantPrice())

	plantType = WALLNUT
	plant, err = CreatePlant(plantType)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Plant Name: ", plant.PlantName())
	fmt.Println("Plant Price: ", plant.PlantPrice())

	// TODO add more plant type ...	

	// Example error case:
	plantType = TypePlant(100)
	plant, err = CreatePlant(plantType)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Plant Name: ", plant.PlantName())
	fmt.Println("Plant Price: ", plant.PlantPrice())
}
