/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Factory Pattern
 */

import Enum.TypePlant;
import Interface.Plant;
import Plants.*;

class PlantFactory {
    public static Plant getPlant(TypePlant type) {
        switch (type) {
            case SUNFLOWER:
                return new SunFlower();
            case PEASHOOTER:
                return new PeaShooter();
            case WALLNUT:
                return new WallNut();
            default:
                System.out.println("Don't have type plant: " + type);
                return null;
        }
    }
}

public class main {
    public static void main(String[] args) {
        Plant sunFlower = PlantFactory.getPlant(TypePlant.SUNFLOWER);
        sunFlower.show();
        System.out.println("====================================");

        Plant peaShooter = PlantFactory.getPlant(TypePlant.PEASHOOTER);
        peaShooter.show();
        System.out.println("====================================");


        Plant wallNut = PlantFactory.getPlant(TypePlant.WALLNUT);
        wallNut.show();
        System.out.println("====================================");

        return;
    }
}