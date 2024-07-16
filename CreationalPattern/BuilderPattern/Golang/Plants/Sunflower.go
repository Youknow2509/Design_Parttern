
package Plants

import (
	"strconv"
	Interface "v/Interface"
)

/**
 * Sunflower struct implements Plant interface
 */
type Sunflower struct {
	Name       string
	Hp         int
	Price      int
	Sun        int
	TimeOutSun int
	SpeedSun   int
}

// Constructor
func NewSunflower(arg ...interface{}) Interface.Plant {
	p := &Sunflower{
		Name:  "SunFlower",
		Hp:    100,
		Price: 50,
	}

	lenArg := len(arg)

	if lenArg > 0 {
		if arg[0] != nil {
			p.Name = arg[0].(string)
		}
		if arg[1] != nil {
			p.Hp = arg[1].(int)
		}
		if arg[2] != nil {
			p.Price = arg[2].(int)
		}
		if arg[3] != nil {
			p.Sun = arg[3].(int)
		}
		if arg[4] != nil {
			p.TimeOutSun = arg[4].(int)
		}
		if arg[5] != nil {
			p.SpeedSun = arg[5].(int)
		}
	}

	return p
}

// Deloy method GetName
func (s *Sunflower) GetName() string {
	return s.Name
}

// Deloy method ToString
func (s *Sunflower) ToString() string {
	return "Name: " + s.Name + ", HP: " + strconv.Itoa(s.Hp) + ", Price: " + strconv.Itoa(s.Price) + ", Speed: " + strconv.Itoa(s.SpeedSun) + ", Sun: " + strconv.Itoa(s.Sun) + ", TimeOutSun: " + strconv.Itoa(s.TimeOutSun)
}

