
package Plants


/**
 * SunFlower struct, implement Plant interface
 */
type SunFlower struct {
	Name  string
	Price int
	// TODO add more attribute ...
}

// Constructor for SunFlower 
func NewSunFlower(params...interface{}) *SunFlower {
	p := &SunFlower {
		Name: "SunFlower",
		Price: 50,
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

func (s *SunFlower) PlantName() string {
	return s.Name
}

func (s *SunFlower) PlantPrice() int {
	return s.Price
}
