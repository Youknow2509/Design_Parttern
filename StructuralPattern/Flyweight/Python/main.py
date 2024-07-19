"""
    @author: Ly Tran Vinh
    @contact: lytranvinh.work@gmail.com
    @content: Flyweight Pattern - Python
"""

from Factory import GameObjectFactory
from Interfaces import GameObject
from Concretes import Plant, Zombie

def main():
    """
        Main function
    """
    factory = GameObjectFactory()
    peashooter = factory.get_Plant("Peashooter", "peashooter.png")
    sunflower = factory.get_Plant("Sunflower", "sunflower.png")
    peashooter2 = factory.get_Plant(type="Peashooter2", image="peashooter2.png")
    basicZombie = factory.get_Plant("BasicZombie", image="basicZombie.png")
    flagZombie = factory.get_Plant("FlagZombie", "flagZombie.png")
    
    peashooter.display(10, 20)
    sunflower.display(30, 40)
    peashooter2.display(50, 60)
    basicZombie.display(70, 80)
    flagZombie.display(90, 100)
    
    if peashooter is peashooter2:
        print("Peashooter and Peashooter2 are the same")
    else:
        print("Peashooter and Peashooter2 are different")    
        
if __name__ == '__main__':
    main()