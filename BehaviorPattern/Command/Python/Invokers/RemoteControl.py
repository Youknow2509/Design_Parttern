
from Interfaces import Command

class RemoteControl:
    """
        Class Remote Control
    """
    def __init__(self):
        self.command: Command = None
        pass
    
    def setCommand(self, command: Command):
        """
            Set command
        """
        self.command = command
        
    def pressButton(self):
        """
            Press button
        """
        self.command.execute()