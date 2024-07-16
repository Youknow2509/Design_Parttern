""" 
    @author: Ly Tran Vinh
	@contact: lytranvinh.work@gmail.com
	@content: Factory Pattern
"""

instance = None

class Person:
    def __init__(self, *args, **kwargs):
        self.name = kwargs.get('name', 'Ly Tran Vinh')
        self.contact = kwargs.get('contact', 'lytranvinh.work@gmail.com')
        if args:
            lenArgs = len(args)
            if lenArgs >= 1:
                self.name = args[0]
            if lenArgs >= 2:
                self.contact = args[1]

    def __str__(self) -> str:
        return 'Name: {0}, Contact: {1}'.format(self.name, self.contact)
    
def getPerson(*args, **kwargs) -> Person:
    global instance
    if not instance:
        instance = Person(*args, **kwargs)
    return instance

def main():
    """
        Main function
    """
    person1 = getPerson()
    print(person1)
    
    person2 = getPerson('dasdasd', 'dasdi@outlook.com')
    print(person2)

if __name__ == '__main__':
    main()