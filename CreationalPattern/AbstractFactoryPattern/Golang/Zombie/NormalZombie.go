
package Zombie


/**
 * NormalZombie struct, implement Zombie interface
 */
type NormalZombie struct {
	Name string
	Dame int
	// TODO add more attribute ...
}

// Constructor for NormalZombie
func NewNormalZombie(params ...interface{}) *NormalZombie {
	p := &NormalZombie{
		Name: "NormalZombie",
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

func (z *NormalZombie) CharacterName() string {
	return z.Name
}

func (z *NormalZombie) GetDamage() int {
	return z.Dame
}
