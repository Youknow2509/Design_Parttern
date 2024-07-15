/**
@author: Ly Tran Vinh
@contact: lytranvinh.work@gmail.com
@content: Abstract Factory Pattern
*/

package main

import (
	"fmt"
)

/**
 * Enum CharacterType
 */
type TypeCharecter int

const (
	PLANT TypeCharecter = iota
	ZOMBIE
)

// String method for CharecterType type
func (d TypeCharecter) String() string {
	return [...]string{
		"Plant",
		"Zombie",
	}[d]
}

// Character interface
type Character interface {
	CharacterName() string
}

/**
 * Enum Type PlantType
 */

type TypePlant int

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
	Character
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
func NewSunFlower(params ...interface{}) *SunFlower {
	p := &SunFlower{
		Name:  "SunFlower",
		Price: 50,
	}

	len := len(params)

	if len > 0 {
		p.Name = params[0].(string)
	}

	if len > 1 {
		p.Price = params[1].(int)
	}

	return p
}

func (s *SunFlower) CharacterName() string {
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
func NewPeaShooter(params ...interface{}) *PeaShooter {
	p := &PeaShooter{
		Name:  "PeaShooter",
		Price: 100,
	}

	len := len(params)

	if len > 0 {
		p.Name = params[0].(string)
	}

	if len > 1 {
		p.Price = params[1].(int)
	}

	return p
}

func (s *PeaShooter) CharacterName() string {
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
func NewWallNut(params ...interface{}) *WallNut {
	p := &WallNut{
		Name:  "WallNut",
		Price: 50,
	}

	len := len(params)

	if len > 0 {
		p.Name = params[0].(string)
	}

	if len > 1 {
		p.Price = params[1].(int)
	}

	return p
}

func (s *WallNut) CharacterName() string {
	return s.Name
}

func (s *WallNut) PlantPrice() int {
	return s.Price
}

/**
 * Zombie
 */
type TypeZombie int

const (
	NORMAL TypeZombie = iota
	CONEHEAD
	// TODO add more zombie type ...
)

// String method for TypeZombie type
func (d TypeZombie) String() string {
	return [...]string{
		"Normal",
		"ConeHead",
		// TODO add more zombie type ...
	}[d]
}

/**
 * Zombie interface
 */
type Zombie interface {
	Character
	GetDamage() int
	// TODO add more method ...
}

/**
 * NormalZombie struct, implement Zombie interface
 */
type NormalZombie struct {
	Name string
	Dame int
	// TODO add more attribute ...
}

// Constructor for NormalZombie
func NewNormalZombie(params ...interface{}) *NormalZombie {
	p := &NormalZombie{
		Name: "NormalZombie",
	}

	len := len(params)

	if len > 0 {
		p.Name = params[0].(string)
	}

	if len > 1 {
		p.Dame = params[1].(int)
	}

	return p
}

func (z *NormalZombie) CharacterName() string {
	return z.Name
}

func (z *NormalZombie) GetDamage() int {
	return z.Dame
}

/**
 * ConeHeadZombie struct, implement Zombie interface
 */
type ConeHeadZombie struct {
	Name string
	Dame int
	// TODO add more attribute ...
}

// Constructor for ConeHeadZombie
func NewConeHeadZombie(params ...interface{}) *ConeHeadZombie {
	p := &ConeHeadZombie{
		Name: "ConeHeadZombie",
	}

	len := len(params)

	if len > 0 {
		p.Name = params[0].(string)
	}

	if len > 1 {
		p.Dame = params[1].(int)
	}

	return p
}

func (z *ConeHeadZombie) CharacterName() string {
	return z.Name
}

func (z *ConeHeadZombie) GetDamage() int {
	return z.Dame
}

/**
 * Struct PlantFactory
 */ 
type PlantFactory struct {}

// Method struct PlantFactory
func (f *PlantFactory) CreateCharacterPlant(typePlant TypePlant) (Character, error) {
	switch typePlant {
	case SUNFLOWER:
		return NewSunFlower(), nil
	case PEASHOOTER:
		return NewPeaShooter(), nil
	case WALLNUT:
		return NewWallNut(), nil
	default:
		return nil, fmt.Errorf("Invalid plant type !!!")
	}

}

/** 
 * Struct ZombieFactory
 */
type ZombieFactory struct {}

// Method struct ZombieFactory
func (f *ZombieFactory) CreateCharacterZombie(typeZombie TypeZombie) (Character, error) {
	switch typeZombie {
	case NORMAL:
		return NewNormalZombie(), nil
	case CONEHEAD:
		return NewConeHeadZombie(), nil
	default:
		return nil, fmt.Errorf("Invalid zombie type !!!")
	}
}

/**
 * Abstract Factory Pattern 
 */
type FactoryProducer struct {}

// Method struct FactoryProducer
func (f *FactoryProducer) GetFactory(typeCharacter TypeCharecter) (interface{}, error) {
	switch typeCharacter {
	case PLANT:
		return &PlantFactory{}, nil
	case ZOMBIE:
		return &ZombieFactory{}, nil
	default:
		return nil, fmt.Errorf("Invalid character type !!!")
	}
}

func main() {
	
	// Create plant
	producer := &FactoryProducer{}
	plantFactory, err := producer.GetFactory(PLANT)
	if err != nil {
		fmt.Println(err)
		return
	}

	plant, err := plantFactory.(*PlantFactory).CreateCharacterPlant(SUNFLOWER)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Plant Name: ", plant.CharacterName())

	plant, err = plantFactory.(*PlantFactory).CreateCharacterPlant(PEASHOOTER)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Plant Name: ", plant.CharacterName())

	plant, err = plantFactory.(*PlantFactory).CreateCharacterPlant(WALLNUT)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Plant Name: ", plant.CharacterName())

	// Create zombie
	zombieFactory, err := producer.GetFactory(ZOMBIE)
	if err != nil {
		fmt.Println(err)
		return
	}

	zombie, err := zombieFactory.(*ZombieFactory).CreateCharacterZombie(NORMAL)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Zombie Name: ", zombie.CharacterName())

	zombie, err = zombieFactory.(*ZombieFactory).CreateCharacterZombie(CONEHEAD)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Zombie Name: ", zombie.CharacterName())
}
