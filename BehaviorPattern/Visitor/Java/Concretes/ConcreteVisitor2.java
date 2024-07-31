package Concretes;

import Interfaces.Visitor;

public class ConcreteVisitor2 implements Visitor {
    @Override
    public void visit(ConcreteElementA element) {
        System.out.println("ConcreteVisitor2: " + element.operationA());
    }

    @Override
    public void visit(ConcreteElementB element) {
        System.out.println("ConcreteVisitor2: " + element.operationB());
    }
}