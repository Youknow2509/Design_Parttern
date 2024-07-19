"""
    @author: Ly Tran Vinh
    @contact: lytranvinh.work@gmail.com
    @content: Facade Pattern
"""

from Facade import GameFacade

def main():
    """
        Main function
    """
    game = GameFacade()
    game.startGame()
    
if __name__ == '__main__':
    main()