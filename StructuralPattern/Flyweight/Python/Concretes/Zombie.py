
from Interfaces import GameObject

class Zombie(GameObject):
    """
        Class Zombie implements GameObject
    """
    def __init__(self, *args, **kwargs):
        self.type = kwargs.get('type', 'type default') # Intrinsic state
        self.img = kwargs.get('img', 'name default')
        
        if args is not None:
            lenArgs = len(args)
            if lenArgs >= 0:
                self.name = args[0]
            if lenArgs >= 1:
                self.img = args[1]
        
    def grow(self):
        print(self.name + " is growing")

    def attackZombie(self):
        print(self.name + " attack zombie")
        
    # Override this display method
    def display(self, x ,y):
        print("Zombie: " + self.type + " with image: " + self.image + " is displayed at (" + str(x) + ", " + str(y) + ")")