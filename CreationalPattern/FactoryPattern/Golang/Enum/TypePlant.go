
package Enum

type TypePlant int

/**
 * Enum Type PlantType
 */

const (
	SUNFLOWER TypePlant = iota
	PEASHOOTER
	WALLNUT
	// TODO add more plant type ...
)

// String method for TypePlant type
func (d TypePlant) String() string {
	return [...]string{
		"SunFlower",
		"PeaShooter",
		"WallNut",
		// TODO add more plant type ...
	}[d]
}