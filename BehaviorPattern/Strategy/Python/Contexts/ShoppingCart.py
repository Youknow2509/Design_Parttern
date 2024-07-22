
class ShoppingCart:
    """
        Context class Shopping Cart
    """
    def __init__(self, *args, **kargs):
        self.paymentStrategy = kargs.get('paymentStrategy', None)
        if args:
            self.paymentStrategy = args[0]
            
    def setPaymentStrategy(self, paymentStrategy):
        self.paymentStrategy = paymentStrategy
        
    def checkout(self, amount):
        self.paymentStrategy.pay(amount)
        