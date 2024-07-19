
package Factory

import (
	"v/Interfaces"
	"v/Concretes"
)

// Flyweight Factory
type GameObjectFactory struct {
    gameObjectPool map[string](iface.GameObject)
}

// Construct a GameObjectFactory
func NewGameObjectFactory() *(GameObjectFactory) {
	return &GameObjectFactory{gameObjectPool: make(map[string](iface.GameObject))}
}

func (gf GameObjectFactory) GetPlant(t string, image string) iface.GameObject {
    key := "Plant-" + t
    check := gf.gameObjectPool[key]
    if check == nil {
        gf.gameObjectPool[key] = Concretes.NewPlant(t, image)
    }
    return gf.gameObjectPool[key]
}

func (gf GameObjectFactory) GetZombie(t string, image string) iface.GameObject {
    key := "Zombie-" + t
    check := gf.gameObjectPool[key]
    if check == nil {
        gf.gameObjectPool[key] = Concretes.NewZombie(t, image)
    }
    return gf.gameObjectPool[key]
}


