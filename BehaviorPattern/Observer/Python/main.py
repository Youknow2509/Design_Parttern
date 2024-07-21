"""
    @author: Ly Tran Vinh
    @contact: lytranvinh.work@gmail.com
    @content: Observer Pattern
"""

from Concretes import *
from Interfaces import *

def main():
    """
        Main function
    """
    zombie = Zombie()
    sunflower = Sunflower()
    peashooter = Peashooter()

    zombie.addObserver(sunflower)        
    zombie.addObserver(peashooter)

    zombie.setAction("Zombie is moving into the lane")
    zombie.setAction("Zombie is eating a plant")

    zombie.removeObserver(sunflower)

    zombie.setAction("Zombie is moving closer to the house")
    
if __name__ == '__main__':
    main()