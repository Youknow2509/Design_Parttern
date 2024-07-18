
from Interfaces import DecoratorPlants, Plant

class IceDecoratorPlant(DecoratorPlants):
    def __init__(self, *args, **kwargs):
        self.name = kwargs.get('name', None)
        self.plant: Plant = kwargs.get('plant', None)
        if args is not None:
            self.plant = args[0]
        super().__init__(*args, **kwargs)

    def __call__(self, cls=None, *args, **kwargs):
        self.__init__(*args, **kwargs)
        if cls is not None:
            cls.attack = self.attack
            cls.get_name = self.get_name
            cls.set_name = self.set_name
        return cls
    
    def get_name(self):
        return self.name
    def set_name(self, name):
        super().set_name(name)
        
    def attack(self) -> str:
        super().attack()
        print(f'\t and slow attack')
        