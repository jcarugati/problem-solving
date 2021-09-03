package work

import (
	"fmt"
)

type Reptile interface {
    Lay() ReptileEgg
}

type ReptileCreator func() Reptile

type ReptileEgg struct {
    CreateReptile ReptileCreator
	Hatched bool    
}

func(egg *ReptileEgg) Hatch() Reptile {
    if egg.Hatched {
		return nil
	}
	egg.Hatched = true
	return egg.CreateReptile()
}

type FireDragon struct {
}

func (f *FireDragon) Lay() ReptileEgg {
	return ReptileEgg{CreateReptile: NewFireDragon, Hatched: false}
}

func NewFireDragon() Reptile {
	return &FireDragon{}
}

func InheritanceTest() {
    var fireDragon FireDragon
    var egg ReptileEgg = fireDragon.Lay()
    var childDragon Reptile = egg.Hatch()
    fmt.Println(childDragon)
	fmt.Println(egg.Hatch())
}
