
package interfaces

// Strategy Interface
type PaymentStrategy interface {
	Pay(amount int)
}