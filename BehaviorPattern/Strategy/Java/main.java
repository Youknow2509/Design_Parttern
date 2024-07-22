/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Strategy Pattern
 */

import Contexts.ShoppingCart;
import Concretes.*; 

public class main {
    public static void main(String[] args) {
        ShoppingCart cart = new ShoppingCart();

        // Paying with credit card
        cart.setPaymentStrategy(new CreditCardStrategy("1234567890123456", "John Doe", "123", "12/23"));
        cart.checkout(100);

        // Paying with PayPal
        cart.setPaymentStrategy(new PayPalStrategy("myemail@example.com", "mypassword"));
        cart.checkout(200);
    }
}