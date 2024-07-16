
package Plant

/**
 * WallNut struct, implement Plant interface
 */
type WallNut struct {
	Name  string
	Price int
	// TODO add more attribute ...
}

// Constructor for WallNut
func NewWallNut(params ...interface{}) *WallNut {
	p := &WallNut{
		Name:  "WallNut",
		Price: 50,
	}

	len := len(params)

	if len > 0 {
		p.Name = params[0].(string)
	}

	if len > 1 {
		p.Price = params[1].(int)
	}

	return p
}

func (s *WallNut) CharacterName() string {
	return s.Name
}

func (s *WallNut) PlantPrice() int {
	return s.Price
}
