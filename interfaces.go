package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (chic Chicken) PrepareDish() {
	fmt.Println("Cooking chicken!")
}

func (sal Salad) PrepareDish() {
	fmt.Println("Chop ingredients")
	fmt.Println("Add Greek dressing")
}

func prepareDishes(dishes []Preparer) {
	fmt.Println("Prepare dishes:")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Println("--Dish: %v--\n", dish)
		dish.PrepareDish()
	}
	fmt.Println("")
}

func main() {
	dishes := []Preparer{Chicken("Chicken Drumsticks"), Salad("Chicken Salad")}
	prepareDishes(dishes)
}
