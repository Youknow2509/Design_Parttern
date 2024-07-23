/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Templete Method Pattern
 */

import Abstract_Templete_Method.Game;
import Concretes.*;

class main {
    public static void main(String[] args) {
         Game game = new Football();
        game.play();
        System.out.println();
        game = new Cricket();
        game.play();
    }
}