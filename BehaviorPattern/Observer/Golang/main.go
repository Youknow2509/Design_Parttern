/**
	@author: Ly Tran Vinh
	@contact: lytranvinh.work@gmail.com
	@content: Observer Pattern
*/
package main

import (
	"v/Concretes"
)

func main()  {
	zombie := concretes.NewZombie().(*concretes.Zombie)
    sunflower := concretes.NewSunflower()
    peashooter := concretes.NewPeashooter()

    zombie.AddObserver(sunflower)        
    zombie.AddObserver(peashooter)

    zombie.SetAction("Zombie is moving into the lane")
    zombie.SetAction("Zombie is eating a plant")

    zombie.RemoveObserver(sunflower)

    zombie.SetAction("Zombie is moving closer to the house")
}