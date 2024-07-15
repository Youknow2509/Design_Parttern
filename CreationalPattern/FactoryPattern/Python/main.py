""" 
    @author: Ly Tran Vinh
	@contact: lytranvinh.work@gmail.com
	@content: Factory Pattern
"""
from enum import Enum


class TypePlant(Enum):
    """ 
        Enum Type Plant
    """
    SUNFLOWER = "SunFlower"
    PEASHOOTER = "PeaShooter"
    WALLNUT = "WallNut"

class Plant:
    """ 
        Plant interface
    """
    def __init__(self, name):
        self.name = name

    def show(self):
        pass

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

class Wall(Plant):
    """
        Wall class
    """
    def __init__(self, *args, **kwargs):
        if args:
            super().__init__(args[0])
            self.price = args[1]
            self.health = args[2]
        elif kwargs:
            super().__init__(kwargs.get('name','WallNut'))
            self.price = kwargs.get('price', 50)
            self.health = kwargs.get('health', 100)
        else:
            super().__init__('WallNut')
            self.price = 50
            self.health = 100
    
    def show(self):
        print("Name: " + self.name 
                + " Price: " + str(self.price) 
                + " Health: " + str(self.health))

class FactoryPlant:
    """
        Factory Plant
    """
    def __init__(self):
        pass

    def create_plant(self, type_plant, *args, **kwargs):
        if type_plant == TypePlant.SUNFLOWER:
            return SunFlower(*args, **kwargs)
        elif type_plant == TypePlant.PEASHOOTER:
            return PeaShooter(*args, **kwargs)
        elif type_plant == TypePlant.WALLNUT:
            return Wall(*args, **kwargs)
        else:
            print("Type plant is not exist")
            return None

def main():
    """
        Main function
    """
    factory = FactoryPlant()

    sunflower = factory.create_plant(TypePlant.SUNFLOWER)
    pea_shooter = factory.create_plant(TypePlant.PEASHOOTER, name='PeaShooter', price=100, health=123, damage=10)
    wall = factory.create_plant(TypePlant.WALLNUT, 'WallNut', 55, 111)
    
    sunflower.show()
    pea_shooter.show()
    wall.show()
    
if __name__ == '__main__':
    main()