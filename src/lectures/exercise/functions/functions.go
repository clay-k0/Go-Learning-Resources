//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//

package main

import "fmt"

//--Requirements:

// * Write a function that accepts a person's name as a function
// * parameter and displays a greeting to that person.
func greetPerson(name string) {
	fmt.Println("Hello,", name)
}

// * Write a function that returns any message, and call it from within
// * fmt.Println()
func randomQuote() string {
	return "Expect the best, but be prepared for the worst."
}

// * Write a function to add 3 numbers together, supplied as arguments, and
// * return the answer
func addThree(x, y, z int) int {
	return x + y + z
}

// * Write a function that returns any number
func randomNum() int {
	return 69
}

// * Write a function that returns any two numbers
func twoRandomNums() (int, int) {
	return 2, 3
}

func main() {
	//* Call every function at least once
	numOne, numTwo := twoRandomNums()
	fmt.Println(numOne, numTwo)

	greetPerson("Clay")

	fmt.Println(randomQuote())

	//* Add three numbers together using any combination of the existing functions.
	//* Print the result
	sum := addThree(randomNum(), numOne, numTwo)
	fmt.Println(sum)

}
