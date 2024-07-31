package Concretes

import (
    "fmt"
    "visitorpattern/Interfaces"
)

// ConcreteVisitor1 is a struct that implements the Visitor interface
type ConcreteVisitor1 struct{}

// Visit implements the visit method for Elements
func (v *ConcreteVisitor1) Visit(element Interfaces.Element) {
    switch e := element.(type) {
    case *ConcreteElementA:
        fmt.Println("ConcreteVisitor1:", e.OperationA())
    case *ConcreteElementB:
        fmt.Println("ConcreteVisitor1:", e.OperationB())
    }
}
