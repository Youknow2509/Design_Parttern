
package Plants

/**
 * PEASHOOTER struct, implement Plant interface
 */
type PeaShooter struct {
	Name  string
	Price int
	// TODO add more attribute ...
}

// Constructor for PeaShooter
func NewPeaShooter(params...interface{}) *PeaShooter {
	p := &PeaShooter {
		Name: "PeaShooter",
		Price: 100,
	}

	len := len(params)

	if (len > 0) {
		p.Name = params[0].(string)
	}

	if (len > 1) {
		p.Price = params[1].(int)
	}

	return p
}

func (s *PeaShooter) PlantName() string {
	return s.Name
}

func (s *PeaShooter) PlantPrice() int {
	return s.Price
}