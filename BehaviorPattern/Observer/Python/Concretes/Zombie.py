from Interfaces import Subject, Observer

class Zombie(Subject): 
    """
        Class Zombie implements the interface Subject
    """
    def __init__(self):
        super().__init__()
        self.observers = []
        self.action = None

    def addObserver(self , observer): 
        self.observers.append(observer)
    
    def removeObserver(self , observer): 
        self.observers.remove(observer)
    
    def notifyObservers(self):
        for observer in self.observers:
            observer.update(self.action)
        
    def setAction(self, action):
        self.action = action
        self.notifyObservers()
    
    def getAction(self) -> str:
        return self.action
    

