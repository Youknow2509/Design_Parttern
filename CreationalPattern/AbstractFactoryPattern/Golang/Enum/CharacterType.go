package enum

/**
 * Enum CharacterType
 */
type TypeCharecter int

const (
	PLANT TypeCharecter = iota
	ZOMBIE
)

// String method for CharecterType type
func (d TypeCharecter) String() string {
	return [...]string{
		"Plant",
		"Zombie",
	}[d]
}
