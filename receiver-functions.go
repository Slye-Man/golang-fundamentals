package main

import "fmt"

type Space struct {
	occupied bool
	//represents a single parling space
}

type ParkingLot struct {
	spaces []Space
	//represents the multiple spaces that we have
}

func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

//function: when a vehicle leaves the parking space
func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {
	//utilizing the functions
	lot := ParkingLot{spaces: make([]Space, 4)}
	fmt.Println("Initial:", lot)
	lot.occupySpace(1)    //using dot notation
	occupySpace(&lot, 3)  //using pointer
	fmt.Println("After occupied:", lot)

	lot.vacateSpace(3)
	fmt.Println("After vacate:", lot)

}
