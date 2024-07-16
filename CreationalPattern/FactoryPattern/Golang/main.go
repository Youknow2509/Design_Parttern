/**
	@author: Ly Tran Vinh
	@contact: lytranvinh.work@gmail.com
	@content: Factory Pattern
*/
package main

import (
	"fmt"
	Enum "v/Enum"
	Interface "v/Interface"
	Plants "v/Plants"
)

/** 
 * Factory method
 */
func CreatePlant(plantType Enum.TypePlant) (Interface.Plant, error) {
	switch plantType {
	case Enum.SUNFLOWER:
		return Plants.NewSunFlower(), nil
	case Enum.PEASHOOTER:
		return Plants.NewPeaShooter(), nil
	case Enum.WALLNUT:
		return Plants.NewWallNut(), nil
	// TODO add more plant type ...
	default:
		return nil, fmt.Errorf("Invalid plant type")
	}
}

func main() {

	plantType := Enum.SUNFLOWER
	plant, err := CreatePlant(plantType)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Plant Name: ", plant.PlantName())
	fmt.Println("Plant Price: ", plant.PlantPrice())

	plantType = Enum.PEASHOOTER
	plant, err = CreatePlant(plantType)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Plant Name: ", plant.PlantName())
	fmt.Println("Plant Price: ", plant.PlantPrice())

	plantType = Enum.WALLNUT
	plant, err = CreatePlant(plantType)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Plant Name: ", plant.PlantName())
	fmt.Println("Plant Price: ", plant.PlantPrice())

	// TODO add more plant type ...	

	// Example error case:
	plantType = Enum.TypePlant(100)
	plant, err = CreatePlant(plantType)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Plant Name: ", plant.PlantName())
	fmt.Println("Plant Price: ", plant.PlantPrice())
}
