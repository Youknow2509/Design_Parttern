from Interfaces import Visitor
from Concretes import ConcreteElementA
from Concretes import ConcreteElementB

class ConcreteVisitor2(Visitor):
    def visit_concrete_element_a(self, element: ConcreteElementA):
        print(f"ConcreteVisitor2: {element.operation_a()}")

    def visit_concrete_element_b(self, element: ConcreteElementB):
        print(f"ConcreteVisitor2: {element.operation_b()}")
