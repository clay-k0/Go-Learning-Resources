//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price int
}

func printInfo(list [4]Product) {
	var totalCost, totalItems int
	for i := 0; i < len(list); i++ {
		listItem := list[i]
		totalCost += listItem.price

		if listItem.name != "" {
			totalItems += 1
		}
	}
	fmt.Println("Last Item:", list[totalItems-1])
	fmt.Println("Total Items:", totalItems)
	fmt.Println("Total Price: $", totalCost)
}

func main() {

	list := [4]Product{
		{name: "bread", price: 10},
		{name: "lettuce", price: 3},
		{name: "ice cream", price: 5},
	}

	fmt.Println("Current List:")
	printInfo(list)

	fmt.Println("\nAdding item to list...")

	list[3] = Product{"pudding", 2}

	fmt.Println("\nNew List:")
	printInfo(list)
}
