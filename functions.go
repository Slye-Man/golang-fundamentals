//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func greetPerson(name string) {
	fmt.Println("Hi there", name, "!") 
}
//* Write a function that returns any message, and call it from within
//  fmt.Println()
	func helloThere() string {
		return "Hello There!"
	}
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
	func addThreeNumbers(num1, num2, num3 int) int {
		return num1 + num2 + num3
	}

//* Write a function that returns any number
	func returnRandomNumber() int {
		return 11
	}

//* Write a function that returns any two numbers
	func returnTwoNumbers() (int,int) {
		return 7, 5
	}

func main() {
	//calling the function
	greetPerson("Bob")
	fmt.Println(helloThere())

	//* Add three numbers together using any combination of the existing functions.
	a, b := returnTwoNumbers()
	answer := addThreeNumbers(returnRandomNumber(), a, b)
//  * Print the result
	fmt.Println(answer)
//* Call every function at least once
}
