/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Facade Pattern
 */

package main

import (
	"v/Facade"
)

func main() {
	facade := facade.NewGameFacade()
	facade.StartGame()
}