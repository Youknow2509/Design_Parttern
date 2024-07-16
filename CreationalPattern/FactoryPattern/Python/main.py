""" 
    @author: Ly Tran Vinh
	@contact: lytranvinh.work@gmail.com
	@content: Factory Pattern
"""

from Enum import TypePlant
from Plants import SunFlower, PeaShooter, Wall

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