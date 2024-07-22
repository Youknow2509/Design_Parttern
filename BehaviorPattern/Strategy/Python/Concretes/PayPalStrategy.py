
from Interfaces import PaymentStrategy

class PayPalStrategy(PaymentStrategy):
    """
        Class PayPalStrategy implementation PaymentStrategy
    """
    def __init__(self, *args, **kwargs):
        self.emailId = kwargs.get('emailId', None)
        self.password = kwargs.get('password', None)
        
        if args:
            lenArgs = len(args)
            if lenArgs > 0:
                self.emailId = args[0]
            if lenArgs > 1:
                self.password = args[1]
        super().__init__(args, kwargs)
        
    # Override Pay 
    def pay(self, amount):
        print("Paid with PayPal: ", amount)