from Interfaces.SupportHandler import SupportHandler
from .SupportRequest import SupportRequest

class Level2SupportHandler(SupportHandler):
    def handle_request(self, request: SupportRequest):
        if request.get_complexity() <= 2:
            print(f"Level 2 support handled the request: {request.get_message()}")
        elif self.next_handler is not None:
            self.next_handler.handle_request(request)
