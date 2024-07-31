package Concretes;

import Interfaces.Visitor;

public class ConcreteVisitor1 implements Visitor {
    @Override
    public void visit(ConcreteElementA element) {
        System.out.println("ConcreteVisitor1: " + element.operationA());
    }

    @Override
    public void visit(ConcreteElementB element) {
        System.out.println("ConcreteVisitor1: " + element.operationB());
    }
}