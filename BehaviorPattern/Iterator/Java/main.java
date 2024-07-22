/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Iterator Pattern
 */

import Concretes.NameCollection;
import Interfaces.Iterator;

public class main {
    public static void main(String[] args) {
        NameCollection nameCollection = new NameCollection();
        Iterator iterator = nameCollection.createIterator();

        while (iterator.hasNext()) {
            String name = (String) iterator.next();
            System.out.println("Name: " + name);
        }
    }
}