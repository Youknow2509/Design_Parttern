package Shape;

import Color.Color;

/**
 * Circle
 */
public class Circle implements Shape {
    
    private Color color;

    // Constructors
    public Circle(Color color) {
        this.color = color;
    }
    
    @Override
    public void drawShape() {
        System.out.println("Circle " + this.color.drawColor());
    }
}

