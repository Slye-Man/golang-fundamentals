package main

import "fmt"

const (
	Small = iota
	Medium
	Large
)

const (
	Ground = iota
	Air
)

type BeltSize int
type Shipping int

func (belt BeltSize) String() string {
	return []string{"Small", "Medium", "Large"}[belt]
}

func (ship Shipping) String() string {
	return []string{"Ground", "Air"}[ship]
}

type Conveyor interface {
	Convey() BeltSize
}

type Shipper interface {
	Ship() Shipping
}

type WareHouseAutomator interface {
	Conveyor
	Shipper
}

type SpamMail struct {
	amount int
}

func (spam SpamMail) String() string {
	return "Spam mail"
}

func (spam *SpamMail) Ship() Shipping {
	return Air
}

func (spam *SpamMail) Convey() BeltSize {
	return Small
}

func automate(item WareHouseAutomator) {
	fmt.Printf("Convey %v on %v conveyor\n", item, item.Convey())
	fmt.Printf("Ship %v via %v\n", item, item.Ship())
}

type ToxicWaste struct {
	weight int
}

func (toxic *ToxicWaste) Sjip() Shipping {
	return Ground
}

func main() {
	mail := SpamMail{30000}
	automate(&mail)

	// waste := ToxicWaste{250}
	// automate(&waste)
}
