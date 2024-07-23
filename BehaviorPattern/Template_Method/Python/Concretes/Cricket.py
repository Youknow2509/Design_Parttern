
from Abstract_Template_Methods import Game 

class Cricket(Game):
    def initialize(self):
        print("Cricket Game Initialized! Start playing.")

    def start_play(self):
        print("Cricket Game Started. Enjoy the game!")

    def end_play(self):
        print("Cricket Game Finished!")