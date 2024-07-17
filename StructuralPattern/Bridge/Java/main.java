/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Bridge Pattern
 */

import Shape.*;
import Color.*;

public class main {
    public static void main(String[] args) {
        
        Shape circle = new Circle(new Red());
        Shape square = new Square(new Blue());
        Shape circle2 = new Circle(new Blue());

        circle.drawShape();
        square.drawShape();
        circle2.drawShape();

        return;
    }
}