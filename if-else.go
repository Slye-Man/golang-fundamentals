package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 7, 11, 23

	if quiz1 > quiz2 {
		fmt.Println("quiz1 scored higher than quiz2")
	} else if quiz1 < quiz2 {
		fmt.Println("quiz2 scored higher than quiz1")
	} else {
		fmt.Println("quiz 1 and quiz 2 scored the same score")
	}

	if average(quiz1, quiz2, quiz3) > 13 {
		fmt.Println("acceptables grades")
	} else {
		fmt.Println("Instructor did a poor job")
	}
}
