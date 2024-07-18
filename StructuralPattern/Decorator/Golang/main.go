/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Decorator Pattern
 */

package main

import (
	c "v/Concretes"
)

func main() {
	peashooter := c.NewPeashooter()
	peashooter.Attack()

	icePeashooter := c.NewIcePlant(peashooter)
	icePeashooter.Attack()
}