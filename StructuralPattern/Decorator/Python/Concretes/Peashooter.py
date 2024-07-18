
from Interfaces import Plant

class Peashooter(Plant):
    """
        Peashooter class implementation for plant
    """
    def __init__(self, *args, **kwargs):
        self.name = kwargs.get('name', 'Peashooter')
        self.damage = kwargs.get('damage', 20)
        super().__init__(*args, **kwargs)

    def get_name(self) -> str:
        return self.name
    
    def set_name(self, name):
        self.name = name
    
    def attack(self) -> str:
        print(f'{self.name} attack dame: {self.damage}')