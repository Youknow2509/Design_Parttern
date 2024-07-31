from Interfaces import Element

class ConcreteElementB(Element):
    def accept(self, visitor):
        visitor.visit_concrete_element_a(self)

    def operation_b(self):
        return "ElementB"
