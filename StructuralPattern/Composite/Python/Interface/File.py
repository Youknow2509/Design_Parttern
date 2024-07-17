
from Concrete import FileSystemComponent

class File(FileSystemComponent):
    """
        Class File
    """
    
    def __init__(self, name):
        """
            Constructor
        """
        super().__init__(name)
    
    def getDetails(self):
        """
            Show all the details
        """
        print("File: " + self.name)