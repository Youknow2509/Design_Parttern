
from Interfaces import GameObject
from Concretes import Plant, Zombie

class GameObjectFactory:
    """
        Factory for creating GameObjects
    """
    def __init__(self, *args, **kwargs):
        self.gameObjectPool = {} # HashMap: Map <String, GameObject>
        pass
    
    def get_Plant(self, type, image) -> GameObject:
        key = "Plant-" + type
        if key not in self.gameObjectPool:
            self.gameObjectPool[key] = Plant(type, image)
        return self.gameObjectPool[key]
    
    def get_Zombie(self, type, image) -> GameObject:
        key = "Zombie-" + type
        if key not in self.gameObjectPool:
            self.gameObjectPool[key] = Zombie(type, image)
        return self.gameObjectPool[key]