/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Vistor Pattern
 */

package main

import (
	"visitorpattern/Concretes"
	"visitorpattern/Interfaces"
)

func main() {
	elements := []Interfaces.Element{&Concretes.ConcreteElementA{}, &Concretes.ConcreteElementB{}}

	visitor1 := &Concretes.ConcreteVisitor1{}
	visitor2 := &Concretes.ConcreteVisitor2{}

	for _, element := range elements {
		element.Accept(visitor1)
		element.Accept(visitor2)
	}
}
