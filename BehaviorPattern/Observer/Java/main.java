/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Observer Pattern
 */
import Interfaces.*;
import Concretes.*;

public class main {
    public static void main(String[] args) {
        Zombie zombie = new Zombie();

        Observer sunflower = new Sunflower();
        Observer peashooter = new Peashooter();

        zombie.addObserver(sunflower);
        zombie.addObserver(peashooter);

        zombie.setAction("Zombie is moving into the lane");
        zombie.setAction("Zombie is eating a plant");

        zombie.removeObserver(sunflower);

        zombie.setAction("Zombie is moving closer to the house");
    }
}
