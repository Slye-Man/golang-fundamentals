//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Coordinate struct{
	x, y int
}

//* Create a rectangle structure containing its coordinates
type Recatangle struct {
	a Coordinate //top left
	b Coordinate // bottom right
}

func width(rect Recatangle) int {
	return (rect.b.x - rect.a.x)
}

func length(rect Recatangle) int {
	return (rect.a.y - rect.b.y)
}
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
func area (rect Recatangle) int {
	return length(rect) * width(rect)
}
func perimeter(rect Recatangle) int {
	return (width(rect) * 2) + (length(rect) * 2)
}
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
func printData(rect Recatangle) {
	fmt.Println("Area is", area(rect))
	fmt.Println("Perimeter is", perimeter(rect))
}

func main() {
	rect := Recatangle{a: Coordinate{4, 7}, b: Coordinate{3, 0}}
	printData(rect)

	rect.a.y *= 2
	rect.b.y *= 2
	printData(rect)
}
