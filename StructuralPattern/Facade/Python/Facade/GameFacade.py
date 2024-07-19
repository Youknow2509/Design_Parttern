
from SubSystems import Plant, Zombie, GameEngine

class GameFacade:
    """
        Class facade for game
    """
    def __init__(self):
        self.gameEngine = GameEngine()
        self.plant = Plant()
        self.zombie = Zombie()
    
    def startGame(self):
        """
            Start game
        """
        self.gameEngine.startGame()
        self.plant.grow()
        self.zombie.spawn()
        self.plant.attackZombie()
        self.gameEngine.endGame()