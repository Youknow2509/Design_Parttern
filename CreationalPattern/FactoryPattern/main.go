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
	SUNFLOWER 	TypePlant = iota 
	PEASHOOTER
	WALLNUT
	REAPEATER
	// TODO add more plant type ...
)

// String method for TypePlant type
func (d TypePlant) String() string {
    return [...]string{
		"SunFlower", 
		"PeaShooter", 
		"WallNut", 
		"ReaPeater",
		// TODO add more plant type ...
	}[d]
}



func main()  {

	fmt.Println(SUNFLOWER)
	fmt.Println(PEASHOOTER)

}