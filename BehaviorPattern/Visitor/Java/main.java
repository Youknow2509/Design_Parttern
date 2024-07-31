import java.util.Arrays;
import java.util.List;

import Concretes.ConcreteElementA;
import Concretes.ConcreteElementB;
import Concretes.ConcreteVisitor1;
import Concretes.ConcreteVisitor2;
import Interfaces.Visitor;

/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Vistor Pattern
 */

import Interfaces.*;

public class main {
    public static void main(String[] args) {
        List<Element> elements = Arrays.asList(new ConcreteElementA(), new ConcreteElementB());

        Visitor visitor1 = new ConcreteVisitor1();
        Visitor visitor2 = new ConcreteVisitor2();

        for (Element element : elements) {
            element.accept(visitor1);
            element.accept(visitor2);
        }
    }
}