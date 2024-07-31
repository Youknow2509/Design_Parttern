package Interfaces

// Element interface defines the accept method that each concrete element will implement
type Element interface {
    Accept(visitor Visitor)
}
