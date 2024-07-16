/**
@author: Ly Tran Vinh
@contact: lytranvinh.work@gmail.com
@content: Prototype Pattern
*/

package main

import (
	"fmt"
	Plants "v/Plants"
)


func main() {

	// Create Sunflower
	sunFlower := Plants.NewSunFlower()
	stringSunFlower := sunFlower.ToString()
	fmt.Println(stringSunFlower)
	fmt.Println("Location memory of Sunflower: ", &sunFlower)

	// Create Clone Sunflower
	cloneSunFlower := sunFlower.Clone()
	// Change properties of Clone Sunflower
	if cloneSunFlower, ok := cloneSunFlower.(*(Plants.Sunflower)); ok {
		cloneSunFlower.Name = "Clone Sunflower"
		cloneSunFlower.Price = 2000
		cloneSunFlower.Health = 300
	} 

	stringCloneSunFlower := cloneSunFlower.ToString()
	fmt.Println(stringCloneSunFlower)
	fmt.Println("Location memory of Clone Sunflower: ", &cloneSunFlower)

}
