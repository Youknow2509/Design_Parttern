
from abc import abstractmethod, ABC

class Iterator(ABC):
    """
        Iterator interface
    """
    def __init__(self, *args, **kwargs):
        pass
    
    @abstractmethod
    def hasNext(self) -> bool:
        """
            Check have next element
        """
        
    @abstractmethod
    def next(self) -> object:
        """
            Get next element
        """