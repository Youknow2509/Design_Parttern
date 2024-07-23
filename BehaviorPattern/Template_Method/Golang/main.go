/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Templete Method Pattern
 */

package main

import (
	"fmt"
	"v/Concretes"
	abs "v/Abstract_Template_Methods"
)

func main() {
	game := concretes.NewFootball()
	abs.Play(game)
	
	fmt.Println()

	game = concretes.NewCricket()
	abs.Play(game)

}