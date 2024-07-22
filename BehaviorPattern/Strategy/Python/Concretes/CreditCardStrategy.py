
from Interfaces import PaymentStrategy

class CreditCardStrategy(PaymentStrategy):
    """
        Class CreditCardStrategy implementation PaymentStrategy
    """
    def __init__(self, *args, **kwargs):
        self.cardNumber = kwargs.get('cardNumber', None)
        self.name = kwargs.get('name', None)
        self.cvv = kwargs.get('cvv', None)
        self.dateOfExpiry = kwargs.get('dateOfExpiry', None)
        
        if args:
            lenArgs = len(args)
            if lenArgs > 0:
                self.cardNumber = args[0]
            if lenArgs > 1:
                self.name = args[1]
            if lenArgs > 2:
                self.cvv = args[2]
            if lenArgs > 3:
                self.dateOfExpiry = args[3]
        super().__init__(args, kwargs)
        
    # Override Pay 
    def pay(self, amount):
        print("Paid with Credit Card: ", amount)