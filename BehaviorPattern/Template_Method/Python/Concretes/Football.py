
from Abstract_Template_Methods import Game 

class Football(Game):
    def initialize(self):
        print("Football Game Initialized! Start playing.")

    def start_play(self):
        print("Football Game Started. Enjoy the game!")

    def end_play(self):
        print("Football Game Finished!")