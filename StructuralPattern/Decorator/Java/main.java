/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Decorator Pattern
 */

import Interfaces.Plant;
import Plants.*;

public class main {
    public static void main(String[] args) {
        Plant basicPlant = new Peashooter();
        basicPlant.attack();
        System.out.println("\n-----");

        Plant firePlant = new FirePlantDecorator(new Peashooter());
        firePlant.attack();
        System.out.println("\n-----");

        Plant iceFirePlant = new IcePlantDecorator(new FirePlantDecorator(new Peashooter()));
        iceFirePlant.attack();
    }
}