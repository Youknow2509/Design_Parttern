package Concretes

import (
    "visitorpattern/Interfaces"
)

// ConcreteElementA is a struct that implements the Element interface
type ConcreteElementA struct{}

// Accept allows the visitor to visit this element
func (e *ConcreteElementA) Accept(visitor Interfaces.Visitor) {
    visitor.Visit(e)
}

// OperationA is a specific method for ConcreteElementA
func (e *ConcreteElementA) OperationA() string {
    return "ElementA"
}
