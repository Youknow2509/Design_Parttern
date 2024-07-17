"""
    @author: Ly Tran Vinh
    @contact: lytranvinh.work@gmail.com
    @content: Bridge Pattern
"""

from Shape import *
from Color import *

def main():
    """
        Function Main 
    """

    circle = Circle(RedColor())
    circle2 = Circle(BlueColor())
    square = Square(GreenColor())
    
    print(circle)
    print(circle2)
    print(square)

if __name__ == '__main__':
    main()