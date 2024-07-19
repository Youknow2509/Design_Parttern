
package Factory;

import Interfaces.GameObject;
import Concretes.Plant;
import Concretes.Zombie;
import java.util.HashMap;
import java.util.Map;

// Flyweight Factory
public class GameObjectFactory {
    private final Map<String, GameObject> gameObjectPool = new HashMap<>();

    public GameObject getPlant(String type, String image) {
        String key = "Plant-" + type;
        GameObject gameObject = gameObjectPool.get(key);
        if (gameObject == null) {
            gameObject = new Plant(type, image);
            gameObjectPool.put(key, gameObject);
        }
        return gameObject;
    }

    public GameObject getZombie(String type, String image) {
        String key = "Zombie-" + type;
        GameObject gameObject = gameObjectPool.get(key);
        if (gameObject == null) {
            gameObject = new Zombie(type, image);
            gameObjectPool.put(key, gameObject);
        }
        return gameObject;
    }
}
