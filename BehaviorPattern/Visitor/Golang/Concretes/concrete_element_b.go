package Concretes

import (
    "visitorpattern/Interfaces"
)

// ConcreteElementB is a struct that implements the Element interface
type ConcreteElementB struct{}

// Accept allows the visitor to visit this element
func (e *ConcreteElementB) Accept(visitor Interfaces.Visitor) {
    visitor.Visit(e)
}

// OperationB is a specific method for ConcreteElementB
func (e *ConcreteElementB) OperationB() string {
    return "ElementB"
}
