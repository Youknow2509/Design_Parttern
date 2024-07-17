
from Concrete import FileSystemComponent

class Folder(FileSystemComponent):
    """
        Class Folder
    """
    def __init__(self, name):
        """
            Constructor
        """
        self.name = name
        self.files = []
        super().__init__(name)
    
    def addFile(self, file):
        """
            Add file
        """
        self.files.append(file)
        
    def removeFile(self, file):
        """
            Remove file
        """
        self.files.remove(file)
    
    def getDetails(self):
        """
            Show all the details
        """
        print("Folder: " + self.name)
        for file in self.files:
            file.getDetails()