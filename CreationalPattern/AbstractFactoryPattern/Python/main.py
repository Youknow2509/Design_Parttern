"""
    @author: Ly Tran Vinh
    @contact: lytranvinh.work@gmail.com
    @content: Abstract Factory Pattern
"""

from enum import Enum

class CharacterType(Enum):
    """
        Enum class to define the type of character
    """
    PLANT = "Plant"
    ZOMBIE = "Zombie"

class PlantType(Enum):
    """
        Enum class to define the type of plant
    """
    SUNFLOWER = "SunFlower"
    PEASHOOTER = "PeaShooter"

class ZombieType(Enum):
    """
        Enum class to define the type of zombie
    """
    NORMAL_ZOMBIE = "NormalZombie"
    FLAG_ZOMBIE = "FlagZombie"

# TODO cmp

def main():
    """
        Main function to run the code
    """


if __name__ == '__main__':
    main()