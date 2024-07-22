
from Interfaces import Command
from Receivers import Light

class LightOffCommand(Command):
    """
        Class Light Off Command implementation Command Interface
    """
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.light: Light = kwargs.get('light', None)
        if args:
            self.light = args[0]
            
    # Override
    def execute(self):
        """
            Execute command
        """
        self.light.off()