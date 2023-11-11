package main

import "fmt"

func double(x int) int {
	return x + x
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello from the greet func!")
}

func main() {
	greet()

	oneFoot := double(6)
	fmt.Println("1 foot is", oneFoot, "inches")

	oneYard := add(oneFoot, 24)
	fmt.Println("1 yard is", oneYard, "inches")

	anotherYard := add(double(6), 24)
	fmt.Println("Once again, 1 yard is", anotherYard, "inches")
}
