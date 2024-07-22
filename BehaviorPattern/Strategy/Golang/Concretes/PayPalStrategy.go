package concretes

import (
	"fmt"
	iface "v/Interfaces"
)

// Concrete Strategy for Credit Card Payment
type PayPalStrategy struct {
	emailId string
	password string
}

// Constructor
func NewPayPalStrategy(emailId string, password string) iface.PaymentStrategy {
	return &PayPalStrategy{
		emailId: emailId,
		password: password,
	}
}

// Override pay
func (c *PayPalStrategy) Pay(amount int) {
	fmt.Println(amount, " paid with pay pal.")
}