/**
@author: Ly Tran Vinh
@contact: lytranvinh.work@gmail.com
@content: Builder Pattern
*/

package main

import (
	"fmt"
	Interface "v/Interface"
	Plants "v/Plants"
)

/**
 * Plant Builder Interface
 */
type PlantBuilder interface {
	Build() Interface.Plant
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
func (s *SunflowerBuilder) Build() Interface.Plant {
	p := &Plants.Sunflower{
		Name:       "SunFlower",
		Hp:         100,
		Price:      50,
		Sun:        50,
		TimeOutSun: 10,
		SpeedSun:   1,
	}

	if s.name != "" {
		p.Name = s.name
	}

	if s.hp != 0 {
		p.Hp = s.hp
	}

	if s.price != 0 {
		p.Price = s.price
	}

	if s.sun != 0 {
		p.Sun = s.sun
	}

	if s.timeOutSun != 0 {
		p.TimeOutSun = s.timeOutSun
	}

	if s.speedSun != 0 {
		p.SpeedSun = s.speedSun
	}

	return p
}

func main() {
	// Use Builder Pattern
	builder := &SunflowerBuilder{}
	plant := builder.SetName("Sunflower 1").SetHp(100).SetPrice(50).SetSun(50).SetTimeOutSun(10).SetSpeedSun(1).Build()
	fmt.Println(plant.ToString())

	builder2 := &SunflowerBuilder{}
	plant2 := builder2.SetName("Sunflower 2").SetHp(100).SetPrice(55).SetTimeOutSun(5).Build()
	fmt.Println(plant2.ToString())

	builder3 := &SunflowerBuilder{}
	plant3 := builder3.Build()
	fmt.Println(plant3.ToString())
}
