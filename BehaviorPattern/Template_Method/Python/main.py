"""
    @author: Ly Tran Vinh
    @contact: lytranvinh.work@gmail.com
    @content: Templete Method Pattern
"""

from Abstract_Template_Methods import Game
from Concretes import *

def main():
    """Main Function
    """
    game = Football()
    game.play()
    print()
    game = Cricket()
    game.play()

if __name__ == "__main__":
    main()