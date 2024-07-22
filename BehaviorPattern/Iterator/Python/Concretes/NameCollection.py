
from Interfaces import *



class NameCollection(Aggregate):
    """
        Class NameCollection implementation Aggregate Interface
    """
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.names = ["Alice", "Bob", "Charlie", "Diana", "Elisa"]
        
    # Override createIterator
    def createIterator(self) -> Iterator:
        """
            Create iterator
        """
        return self.NameIterator(self, self.names)
    
    class NameIterator(Iterator):
        """
            Class NameIterator implementation Iterator Interface
        """
        def __init__(self, *args, **kwargs):
            super().__init__(*args, **kwargs)
            self.names = args[1]
            self.index = 0
            
        def hasNext(self) -> bool:
            """
                Check have next element
            """
            return self.index < len(self.names)
                
        def next(self) -> object:
            """
                Get next element
            """
            if self.hasNext():
                name = self.names[self.index]
                self.index += 1
                return name
            return None