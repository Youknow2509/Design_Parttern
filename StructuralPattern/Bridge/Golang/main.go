/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Bridge Pattern
 */

package main

import (
	cl "v/color"
	sh "v/shape"
)

func main() {

	circle := sh.NewCircle(&cl.BlueColor{})
	circle.Draw()

	square := sh.NewSquare(&cl.RedColor{})
	square.Draw()

}
