package Concretes

import (
    "fmt"
    "visitorpattern/Interfaces"
)

// ConcreteVisitor2 is a struct that implements the Visitor interface
type ConcreteVisitor2 struct{}

// Visit implements the visit method for Elements
func (v *ConcreteVisitor2) Visit(element Interfaces.Element) {
    switch e := element.(type) {
    case *ConcreteElementA:
        fmt.Println("ConcreteVisitor2:", e.OperationA())
    case *ConcreteElementB:
        fmt.Println("ConcreteVisitor2:", e.OperationB())
    }
}
