
package Facade;

import SubSystems.*;

public class GameFacade {
    private Plant plant;
    private Zombie zombie;
    private GameEngine gameEngine;

    public GameFacade() {
        this.plant = new Plant();
        this.zombie = new Zombie();
        this.gameEngine = new GameEngine();
    }

    public void playGame() {
        gameEngine.startGame();
        plant.grow();
        zombie.spawn();
        plant.attackZombie();
        zombie.attackPlant();
        gameEngine.endGame();
    }
}