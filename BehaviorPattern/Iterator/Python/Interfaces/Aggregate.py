
from abc import abstractmethod, ABC
from .Iterator import Iterator

class Aggregate(ABC):
    """
        Interface Aggregate
    """
    def __init__(self, *args, **kwargs):
        pass
    
    @abstractmethod
    def createIterator() -> Iterator:
        """
            Create iterator
        """
        pass