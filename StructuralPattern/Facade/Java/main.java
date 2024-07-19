/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Facade Pattern
 */

import Facade.GameFacade;

public class main {
    public static void main(String[] args) {
        GameFacade gameFacade = new GameFacade();
        gameFacade.playGame();
    }
}