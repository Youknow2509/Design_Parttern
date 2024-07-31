"""
    @author: Ly Tran Vinh
    @contact: lytranvinh.work@gmail.com
    @content: Visitor Pattern
"""
from Concretes import ConcreteElementA
from Concretes import ConcreteElementB
from Concretes import ConcreteVisitor1
from Concretes import ConcreteVisitor2

def main():
    elements = [ConcreteElementA(), ConcreteElementB()]

    visitor1 = ConcreteVisitor1()
    visitor2 = ConcreteVisitor2()

    for element in elements:
        element.accept(visitor1)
        element.accept(visitor2)

if __name__ == "__main__":
    main()
