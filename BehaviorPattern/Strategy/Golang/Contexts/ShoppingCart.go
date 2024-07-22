
package contexts

import (
	iface "v/Interfaces"
)

// Context Class
type ShoppingCart struct {
	paymentStrategy iface.PaymentStrategy
}

// Constructor 
func NewShoppingCart(paymentStrategy iface.PaymentStrategy) *ShoppingCart {
	return &ShoppingCart{paymentStrategy: paymentStrategy}
}

func (c *ShoppingCart) SetPaymentStrategy(paymentStrategy iface.PaymentStrategy) {
	c.paymentStrategy = paymentStrategy
}

func (c *ShoppingCart) Checkout(amount int) {
	c.paymentStrategy.Pay(amount);
}