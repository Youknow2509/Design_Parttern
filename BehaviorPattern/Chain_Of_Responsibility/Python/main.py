"""
    @author: Ly Tran Vinh
    @contact: lytranvinh.work@gmail.com
    @content: Chain Of Responsibility Pattern
"""

from Concretes import *

def main():
    """
        Main function
    """
    level1_handler = Level1SupportHandler()
    level2_handler = Level2SupportHandler()
    level3_handler = Level3SupportHandler()

    level1_handler.set_next_handler(level2_handler)
    level2_handler.set_next_handler(level3_handler)

    request1 = SupportRequest("Simple issue", 1)
    request2 = SupportRequest("Moderate issue", 2)
    request3 = SupportRequest("Complex issue", 3)
    request4 = SupportRequest("Very complex issue", 4)

    level1_handler.handle_request(request1)
    level1_handler.handle_request(request2)
    level1_handler.handle_request(request3)
    level1_handler.handle_request(request4)
    
if __name__ == '__main__':
    main()