"""
    @author: Ly Tran Vinh
    @contact: lytranvinh.work@gmail.com
    @content: Adapter Pattern
"""

"""
    Adapter Pattern for translate vietnamese to english
"""

class Vietnamese:
    """
        Interface for Vietnamese
    """
    def __init__(self):
        pass
    
    def send(self, content):
        pass
    
class AdapteePattern:
    """
        Adaptee Pattern - Handle translation vietnamese to english
    """
    def __init__(self):
        pass
    
    def send(self, content):
        """
            Translate vietnamese to english
        """
        # TODO: handle it
        if content=="xin chao":
            content = "hello"
        
        # TODO: I only for example iti, if want: use translation api
        return "Translated: " + content

class AdepterPattern:
    """
        Adapter Pattern
    """
    def __init__(self):
        self.adaptee = AdapteePattern()
    
    def send(self, content):
        """
            Send content
        """
        return self.adaptee.send(content)

def main():
    """
        Main function
    """
    vietnamese = AdepterPattern()
    print(vietnamese.send("xin chao"))

if __name__ == '__main__':
    main()