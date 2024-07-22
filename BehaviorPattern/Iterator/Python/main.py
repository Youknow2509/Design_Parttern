"""
    @author: Ly Tran Vinh
    @contact: lytranvinh.work@gmail.com
    @content: Iterator Pattern
"""
from Concretes import NameCollection

def main():
    """
        Main function
    """
    name_collection = NameCollection()
    iterator = name_collection.createIterator()

    while iterator.hasNext():
        name = iterator.next()
        print("Name:", name)

if __name__ == '__main__':
    main()