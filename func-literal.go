package main

import "fmt"

func add(lhs, rhs int) int {
	return lhs + rhs
}

func compute(lhs, rhs int, operation func(lhs, rhs int) int) int {
	fmt.Printf("Running a computation with %v and %v\n", lhs, rhs)
	return operation(lhs, rhs)
}

func main() {
	fmt.Println("2 + 2 =", compute(2, 2, add))

	fmt.Println("11 - 9 =", compute(11, 9, func(lhs, rhs int) int {
		return lhs - rhs
	}))

	mul := func(lhs, rhs int) int {
		fmt.Printf("Multiplying %v * %v = ", lhs, rhs)
		return lhs * rhs
	}

	fmt.Println(compute(4, 7, mul))
}
