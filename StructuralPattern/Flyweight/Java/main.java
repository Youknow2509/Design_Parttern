/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Flyweight Pattern
 */

import Factory.GameObjectFactory;
import Interfaces.GameObject;

public class main {
    public static void main(String[] args) {
        GameObjectFactory factory = new GameObjectFactory();

        GameObject peashooter = factory.getPlant("Peashooter", "peashooter.png");
        GameObject sunflower = factory.getPlant("Sunflower", "sunflower.png");
        GameObject peashooterAgain = factory.getPlant("Peashooter", "peashooter.png");
        GameObject basicZombie = factory.getZombie("BasicZombie", "basicZombie.png");
        GameObject coneheadZombie = factory.getZombie("ConeheadZombie", "coneheadZombie.png");

        peashooter.display(10, 20);
        sunflower.display(30, 40);
        peashooterAgain.display(50, 60);
        basicZombie.display(70, 80);
        coneheadZombie.display(90, 100);

        System.out.println("Are 'peashooter' and 'peashooterAgain' the same object? " + (peashooter == peashooterAgain));
    }
}
