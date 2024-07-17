package Shape;

import Color.Color;



/**
 * Square
 */
public class Square implements Shape {
    
    private Color color;

    // Constructors
    public Square(Color color) {
        this.color = color;
    }
    
    @Override
    public void drawShape() {
        System.out.println("Square " + this.color.drawColor());
    }
}