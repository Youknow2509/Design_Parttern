from Interfaces.SupportHandler import SupportHandler
from .SupportRequest import SupportRequest

class Level3SupportHandler(SupportHandler):
    def handle_request(self, request: SupportRequest):
        if request.get_complexity() <= 3:
            print(f"Level 3 support handled the request: {request.get_message()}")
        else:
            print("Request is too complex to be handled.")
