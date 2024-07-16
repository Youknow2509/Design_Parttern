
package Zombie


/**
 * ConeHeadZombie struct, implement Zombie interface
 */
type ConeHeadZombie struct {
	Name string
	Dame int
	// TODO add more attribute ...
}

// Constructor for ConeHeadZombie
func NewConeHeadZombie(params ...interface{}) *ConeHeadZombie {
	p := &ConeHeadZombie{
		Name: "ConeHeadZombie",
	}

	len := len(params)

	if len > 0 {
		p.Name = params[0].(string)
	}

	if len > 1 {
		p.Dame = params[1].(int)
	}

	return p
}

func (z *ConeHeadZombie) CharacterName() string {
	return z.Name
}

func (z *ConeHeadZombie) GetDamage() int {
	return z.Dame
}