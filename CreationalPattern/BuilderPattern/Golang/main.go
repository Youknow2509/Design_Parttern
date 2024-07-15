/**
	@author: Ly Tran Vinh
	@contact: lytranvinh.work@gmail.com
	@content: Builder Pattern
*/

package main

import (
	"strconv"
)

/**
 * Plant interface
 */
type Plant interface {
	GetName() string
	ToString() string
}

/**
 * Sunflower struct implements Plant interface
 */
type Sunflower struct {
	name       string
	hp         int
	price      int
	sun        int
	timeOutSun int
	speedSun   int
}

// Constructor
func NewSunflower(arg ...interface{}) Plant {
	p := &Sunflower{
		name:  "SunFlower",
		hp:    100,
		price: 50,
	}

	lenArg := len(arg)

	if lenArg > 0 {
		if arg[0] != nil {
			p.name = arg[0].(string)
		}
		if arg[1] != nil {
			p.hp = arg[1].(int)
		}
		if arg[2] != nil {
			p.price = arg[2].(int)
		}
		if arg[3] != nil {
			p.sun = arg[3].(int)
		}
		if arg[4] != nil {
			p.timeOutSun = arg[4].(int)
		}
		if arg[5] != nil {
			p.speedSun = arg[5].(int)
		}
	}

	return p
}

// Deloy method GetName
func (s *Sunflower) GetName() string {
	return s.name
}

// Deloy method ToString
func (s *Sunflower) ToString() string {
	return "Name: " + s.name + ", HP: " + strconv.Itoa(s.hp) + ", Price: " + strconv.Itoa(s.price) + ", Speed: " + strconv.Itoa(s.speedSun) + ", Sun: " + strconv.Itoa(s.sun) + ", TimeOutSun: " + strconv.Itoa(s.timeOutSun)
}

/**
 * Plant Builder Interface
 */
type PlantBuilder interface {
	Build() Plant
	SetName(name string) PlantBuilder
}

/**
 * Struct Builder Sunflower
 */
type SunflowerBuilder struct {
	name       string
	hp         int
	price      int
	sun        int
	timeOutSun int
	speedSun   int
}

// Deloy method SetName
func (s *SunflowerBuilder) SetName(name string) *SunflowerBuilder {
	s.name = name
	return s
}

// Deloy method SetHp
func (s *SunflowerBuilder) SetHp(hp int) *SunflowerBuilder {
	s.hp = hp
	return s
}

// Deloy method SetPrice
func (s *SunflowerBuilder) SetPrice(price int) *SunflowerBuilder {
	s.price = price
	return s
}

// Deloy method SetSun
func (s *SunflowerBuilder) SetSun(sun int) *SunflowerBuilder {
	s.sun = sun
	return s
}

// Deloy method SetTimeOutSun
func (s *SunflowerBuilder) SetTimeOutSun(timeOutSun int) *SunflowerBuilder {
	s.timeOutSun = timeOutSun
	return s
}

// Deloy method SetSpeedSun
func (s *SunflowerBuilder) SetSpeedSun(speedSun int) *SunflowerBuilder {
	s.speedSun = speedSun
	return s
}

// Deloy method Build
func (s *SunflowerBuilder) Build() Plant {
	p := &Sunflower{
		name:       "SunFlower",
		hp:         100,
		price:      50,
		sun:        50,
		timeOutSun: 10,
		speedSun:   1,
	}

	if s.name != "" {
		p.name = s.name
	}

	if s.hp != 0 {
		p.hp = s.hp
	}

	if s.price != 0 {
		p.price = s.price
	}

	if s.sun != 0 {
		p.sun = s.sun
	}

	if s.timeOutSun != 0 {
		p.timeOutSun = s.timeOutSun
	}

	if s.speedSun != 0 {
		p.speedSun = s.speedSun
	}

	return p
}

func main() {
	// Use Builder Pattern
	builder := &SunflowerBuilder{}
	plant := builder.SetName("Sunflower 1").SetHp(100).SetPrice(50).SetSun(50).SetTimeOutSun(10).SetSpeedSun(1).Build()
	println(plant.ToString())

	builder2 := &SunflowerBuilder{}
	plant2 := builder2.SetName("Sunflower 2").SetHp(100).SetPrice(55).SetTimeOutSun(5).Build()
	println(plant2.ToString())

	builder3 := &SunflowerBuilder{}
	plant3 := builder3.Build()
	println(plant3.ToString())
}
