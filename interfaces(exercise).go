//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Lift int

type PickLift interface {
	LiftPick() Lift
}

type Motorcycle string
type Car string
type Truck string

//* Vehicles have a model name in addition to the vehicle type:
func (motor Motorcycle) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(motor))
}

//* Vehicles have a model name in addition to the vehicle type:
func (ca Car) String() string {
	return fmt.Sprintf("Car: %v", string(ca))
}

//* Vehicles have a model name in addition to the vehicle type:
func (truc Truck) String() string {
	return fmt.Sprintf("Truck: %v", string(truc))
}

func (motor Motorcycle) LiftPick() Lift {
	return SmallLift
}

func (ca Car) LiftPick() Lift {
	return StandardLift
}

func (truc Truck) LiftPick() Lift {
	return LargeLift
}

//* Write a single function to handle all of the vehicles
//  that the shop works on.
func proceedToLift(pick PickLift) {
	switch pick.LiftPick() {
		case SmallLift:
			fmt.Printf("Send %v to small lift\n", pick)
		case StandardLift:
			fmt.Printf("Send %v to standard lift\n", pick)
		case LargeLift:
			fmt.Printf("Send %v to large lift\n", pick)
	}
}

func main() {
	car := Car("Driftstang")
	truck := Truck("MANNNNNN")
	motorcycle := Motorcycle("DarveySaint")
	//* Direct at least 1 of each vehicle type to the correct
	//  lift, and print out the vehicle information.
	proceedToLift(truck)
	proceedToLift(motorcycle)
	proceedToLift(car)
	
}
