
from Interface import Plant

class SunFlower(Plant):
    """ 
        SunFlower class
    """
    def __init__(self, *args, **kwargs):
        if args:
            super().__init__(args[0])
            self.price = args[1]
            self.health = args[2]
        elif kwargs:
            super().__init__(kwargs.get('name','SunFlower'))
            self.price = kwargs.get('price', 50)
            self.health = kwargs.get('health', 100)
        else:
            super().__init__('SunFlower')
            self.price = 50
            self.health = 100
    
    def show(self):
        print("Name: " + self.name 
                + " Price: " + str(self.price) 
                + " Health: " + str(self.health))
