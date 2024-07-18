
from .Plant import Plant

class DecoratorPlants(Plant): 
    """
        Interface decorator for plant implementations plant
    """
    def __init__(self, *args, **kwargs):
        self.plant = kwargs.get("plant", Plant)
        if args is not None:
            self.plant = args[0]
        super().__init__(*args, **kwargs)

    def __call__(self, *args, **kwargs):
        pass
    
    def attack(self):
        self.plant.attack()

    