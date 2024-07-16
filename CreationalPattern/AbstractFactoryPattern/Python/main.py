"""
    @author: Ly Tran Vinh
    @contact: lytranvinh.work@gmail.com
    @content: Abstract Factory Pattern
"""

from Enum import *
from Character import *

# TODO: Create abs factory pattern

def main():
    """
        Main function to run the code
    """
    ps = PeaShooter("test", 123, 123)
    print(ps)

    nz = NormalZombie()
    print(nz)

if __name__ == '__main__':
    main()