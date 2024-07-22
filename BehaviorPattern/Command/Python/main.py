"""
    @author: Ly Tran Vinh
    @contact: lytranvinh.work@gmail.com
    @content: Command Pattern
"""

from Interfaces import Command
from Concretes import *
from Receivers import Light
from Invokers import RemoteControl

def main():
    """
        Main function
    """
    light = Light()

    lightOn = LightOnCommand(light)
    lightOff = LightOffCommand(light)

    remote = RemoteControl()

    #  Turn the light on
    remote.setCommand(lightOn)
    remote.pressButton()

    #  Turn the light off
    remote.setCommand(lightOff)
    remote.pressButton()
    
if __name__ == '__main__':
    main()