package concretes

import (
	"fmt"
	"v/Interfaces"
)

// Concrete Strategy for Credit Card Payment
type CreditCardStrategy struct {
	cardNumber string
	name string
	cvv string
	dateOfExpiry string
}

// Constructor
func NewCreditCardStrategy(cardNumber string, name string, cvv string, dateOfExpiry string) interfaces.PaymentStrategy {
	return &CreditCardStrategy{
		cardNumber: cardNumber,
		name: name,
		cvv: cvv,
		dateOfExpiry: dateOfExpiry,
	}
}

// Override pay
func (c *CreditCardStrategy) Pay(amount int) {
	fmt.Println(amount, " paid with credit card.")
}