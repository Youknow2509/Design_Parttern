"""
    @author: Ly Tran Vinh
    @contact: lytranvinh.work@gmail.com
    @content: Strategy Pattern
"""

from Concretes import *
from Contexts import *
from Interfaces import *

def main():
    """
        Main Function
    """
    cart = ShoppingCart()

    # Paying with credit card
    cart.setPaymentStrategy(CreditCardStrategy("1234567890123456", "John Doe", "123", "12/23"))
    cart.checkout(100)

    # Paying with PayPal
    cart.setPaymentStrategy(PayPalStrategy("myemail@example.com", "mypassword"))
    cart.checkout(200)

if __name__ == '__main__':
    main()