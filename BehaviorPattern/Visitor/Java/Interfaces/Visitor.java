package Interfaces;

import Concretes.ConcreteElementA;
import Concretes.ConcreteElementB;

public interface Visitor {
    void visit(ConcreteElementA element);
    void visit(ConcreteElementB element);
}
