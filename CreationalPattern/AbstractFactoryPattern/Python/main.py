"""
    @author: Ly Tran Vinh
    @contact: lytranvinh.work@gmail.com
    @content: Abstract Factory Pattern
"""

from Enum import *
from Character import *

# TODO: Create abs factory pattern
class CharacterFactory:
    """
        Factory - Interface
    """
    def __init__(self) -> None:
        pass
    def createPlant(self, plantType, *args, **kwargs) -> Plant:
        pass
    def createZombie(self, zombieType, *args, **kwargs) -> Zombie:
        pass

class PlantFactory(CharacterFactory):
    """
        Concrete Factory
    """
    def __init__(self) -> None:
        super().__init__()
    def createZombie(self, zombieType, *args, **kwargs) -> Zombie:
        print("Factory Plant can't create Zombie")
        return None
    def createPlant(self, plantType, *args, **kwargs) -> Plant:
        if plantType == PlantType.PEASHOOTER:
            return PeaShooter(*args, **kwargs)
        elif plantType == PlantType.SUNFLOWER:
            return SunFlower(*args, **kwargs)
        else:
            print("Invalid Plant Type: " + str(plantType))
            return None
    
           
class ZombieFactory(CharacterFactory):
    """
        Concrete Factory
    """
    def __init__(self) -> None:
        super().__init__()
    def createPlant(self, plantType, *args, **kwargs) -> Plant:
        print("Factory Zombie can't create Plant")
        return None
    def createZombie(self, zombieType, *args, **kwargs) -> Zombie:
        if zombieType == ZombieType.NORMAL_ZOMBIE:
            return NormalZombie(*args, **kwargs)
        elif zombieType == ZombieType.FLAG_ZOMBIE:
            return FlagZombie(*args, **kwargs)
        else:
            print("Invalid Zombie Type: " + str(zombieType))
            return None

class FactoryProducer:
    """
        Abstract Factory Producer
    """
    def __init__(self) -> None:
        pass
    def getFactory(self, characterType) -> CharacterFactory:
        if characterType == CharacterType.PLANT:
            return PlantFactory()
        elif characterType == CharacterType.ZOMBIE:
            return ZombieFactory()
        else:
            print("Invalid Character Type: " + str(characterType))
            return None

def main():
    """
        Main function to run the code
    """
    producerFactory = FactoryProducer().getFactory(CharacterType.PLANT)

    sunFlower = producerFactory.createPlant(PlantType.SUNFLOWER, 100, 100)
    print(sunFlower)

    peaShooter = producerFactory.createPlant(PlantType.PEASHOOTER, 222)
    print(peaShooter)

    producerFactory = FactoryProducer().getFactory(CharacterType.ZOMBIE)
    nz = producerFactory.createZombie(ZombieType.NORMAL_ZOMBIE)
    print(nz)

    fz = producerFactory.createZombie(ZombieType.FLAG_ZOMBIE, name="Flag Zombie", health=111, damage=22)
    print(fz)

if __name__ == '__main__':
    main()