/**
@author: Ly Tran Vinh
@contact: lytranvinh.work@gmail.com
@content: Abstract Factory Pattern
*/

package main

import (
	"fmt"
	enum "v/Enum"
	Interface "v/Interface"
	Plant "v/Plant"
	Zombie "v/Zombie"
)

/**
 * Struct PlantFactory
 */ 
type PlantFactory struct {}

// Method struct PlantFactory
func (f *PlantFactory) CreateCharacterPlant(typePlant enum.TypePlant) (Interface.Character, error) {
	switch typePlant {
	case enum.SUNFLOWER:
		return Plant.NewSunFlower(), nil
	case enum.PEASHOOTER:
		return Plant.NewPeaShooter(), nil
	case enum.WALLNUT:
		return Plant.NewWallNut(), nil
	default:
		return nil, fmt.Errorf("Invalid plant type !!!")
	}

}

/** 
 * Struct ZombieFactory
 */
type ZombieFactory struct {}

// Method struct ZombieFactory
func (f *ZombieFactory) CreateCharacterZombie(typeZombie enum.TypeZombie) (Interface.Character, error) {
	switch typeZombie {
	case enum.NORMAL:
		return Zombie.NewNormalZombie(), nil
	case enum.CONEHEAD:
		return Zombie.NewConeHeadZombie(), nil
	default:
		return nil, fmt.Errorf("Invalid zombie type !!!")
	}
}

/**
 * Abstract Factory Pattern 
 */
type FactoryProducer struct {}

// Method struct FactoryProducer
func (f *FactoryProducer) GetFactory(typeCharacter enum.TypeCharecter) (interface{}, error) {
	switch typeCharacter {
	case enum.PLANT:
		return &PlantFactory{}, nil
	case enum.ZOMBIE:
		return &ZombieFactory{}, nil
	default:
		return nil, fmt.Errorf("Invalid character type !!!")
	}
}

func main() {
	
	// Create plant
	producer := &FactoryProducer{}
	plantFactory, err := producer.GetFactory(enum.PLANT)
	if err != nil {
		fmt.Println(err)
		return
	}

	plant, err := plantFactory.(*PlantFactory).CreateCharacterPlant(enum.SUNFLOWER)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Plant Name: ", plant.CharacterName())

	plant, err = plantFactory.(*PlantFactory).CreateCharacterPlant(enum.PEASHOOTER)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Plant Name: ", plant.CharacterName())

	plant, err = plantFactory.(*PlantFactory).CreateCharacterPlant(enum.WALLNUT)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Plant Name: ", plant.CharacterName())

	// Create zombie
	zombieFactory, err := producer.GetFactory(enum.ZOMBIE)
	if err != nil {
		fmt.Println(err)
		return
	}

	zombie, err := zombieFactory.(*ZombieFactory).CreateCharacterZombie(enum.NORMAL)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Zombie Name: ", zombie.CharacterName())

	zombie, err = zombieFactory.(*ZombieFactory).CreateCharacterZombie(enum.CONEHEAD)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Zombie Name: ", zombie.CharacterName())
}
