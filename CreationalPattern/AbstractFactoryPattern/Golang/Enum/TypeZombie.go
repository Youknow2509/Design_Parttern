
package enum

/**
 * Zombie Type enum
 */
type TypeZombie int

const (
	NORMAL TypeZombie = iota
	CONEHEAD
	// TODO add more zombie type ...
)

// String method for TypeZombie type
func (d TypeZombie) String() string {
	return [...]string{
		"Normal",
		"ConeHead",
		// TODO add more zombie type ...
	}[d]
}
