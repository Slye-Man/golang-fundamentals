package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {
	switch pri := price(); {
	case pri < 2:
		fmt.Println("Item: cheap")
	case pri < 10:
		fmt.Println("Itme: moderately cheap")
	default:
	fmt.Println("Itme: expensive!!")
	}

	//ticket := Economy
	ticket := FirstClass
	//ticket := Business
	//ticket := Commercial
	
	switch ticket {
	case Economy:
		fmt.Println("Economy seating")
	case Business:
		fmt.Println("Business seating")
	case FirstClass:
		fmt.Println("First Class seating")
	default:
		fmt.Println("Other seating")
	}
}
