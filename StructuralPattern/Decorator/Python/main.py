"""
    @author: Ly Tran Vinh
    @contact: lytranvinh.work@gmail.com
    @content: Decorator Pattern
"""

from Interfaces import *
from Concretes import *

def main():
    """
        Main function
    """
    peashooter = Peashooter()
    peashooter.attack()
    
    ice_decorator_plant = IceDecoratorPlant(peashooter)
    ice_decorator_plant.attack()
        
if __name__ == '__main__':
    main()