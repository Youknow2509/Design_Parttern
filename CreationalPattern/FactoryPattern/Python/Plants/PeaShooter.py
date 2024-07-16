
from Interface import Plant

class PeaShooter(Plant):
    """
        PeaShooter class
    """
    def __init__(self, *args, **kwargs):
        if args:
            super().__init__(args[0])
            self.price = args[1]
            self.health = args[2]
            self.damage = args[3]
        elif kwargs:
            super().__init__(kwargs.get('name','PeaShooter'))
            self.price = kwargs.get('price', 100)
            self.health = kwargs.get('health', 100)
            self.damage = kwargs.get('damage', 10)
        else:
            super().__init__('PeaShooter')
            self.price = 100
            self.health = 100
            self.damage = 10
    
    def show(self):
        print("Name: " + self.name 
                + " Price: " + str(self.price) 
                + " Health: " + str(self.health) 
                + " Damage: " + str(self.damage))
