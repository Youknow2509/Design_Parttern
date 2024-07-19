/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Flyweight Pattern
 */

package main

import (
	"v/Factory"
)

func main() {
	factory := Factory.NewGameObjectFactory()

	peashooter := factory.GetPlant("Peashooter", "peashooter.png")
	zombie := factory.GetZombie("Zombie", "zombie.png")
	peashooter2 := factory.GetPlant("Peashooter2", "peashooter.png")

	peashooter.Display(1, 2)
	zombie.Display(3, 4)
	peashooter2.Display(5, 6)

	if (peashooter == peashooter2) {
		println("Peashooter and Peashooter2 are the same object")
	} else {
		println("Peashooter and Peashooter2 are different objects")
	}

}