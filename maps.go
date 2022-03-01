package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	// len(shoppingList)
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1
	shoppingList["eggs"] = 5

	shoppingList["eggs"] += 1
	fmt.Println(shoppingList)

	delete(shoppingList, "Milk")
	fmt.Println("Milk is deleted, new list:", shoppingList)

	fmt.Println("need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]
	fmt.Println("Need cereal?")
	if !found {
		fmt.Println("nope... cereal not in stock")
	} else {
		fmt.Println("you're in luck", cereal, "boxes")
	}

	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}
	fmt.Println("There are", totalItems, "items on the shopping list")
}
