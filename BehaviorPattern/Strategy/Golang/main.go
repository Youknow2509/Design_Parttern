/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Strategy Pattern
 */

package main

import (
	"v/Concretes"
	"v/Contexts"
)

func main()  {
	cart := &contexts.ShoppingCart{}

	// Paying with credit card
	cart.SetPaymentStrategy(concretes.NewCreditCardStrategy("1234567890123456", "John Doe", "123", "12/23"));
	cart.Checkout(100);

	// Paying with PayPal
	cart.SetPaymentStrategy(concretes.NewPayPalStrategy("myemail@example.com", "mypassword"));
	cart.Checkout(200);

}