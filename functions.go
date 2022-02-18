package main

import "fmt"

//creating a function that doubles a number
func double(num int) int {
	return num * 2
}

//add function
func add(lhs, rhs int) int {
	return lhs + rhs
}

//greet function that will display a message
//notice there are no paramaters
func greet() {
	fmt.Println(("Hello there from the greet funtion"))
}

func main() {
	//calling the function
	greet()

	dozen := double(6)
	fmt.Println("A dozen is", dozen)

	bakersDozen := add(dozen, 1)
	fmt.Println("A baker's dozen is", bakersDozen)

	anotherBakersDozen := add(double(6), 1)
	fmt.Println("another baker's dozen is", anotherBakersDozen)
}
