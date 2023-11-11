//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

// * Create a rectangle structure containing a length and width field
type Rectangle struct {
	Length int
	Width  int
}

// * Using functions, calculate the area and perimeter of a rectangle
//   - The functions must use the rectangle structure as the function parameter
func calculateArea(rectangle Rectangle) int {
	return rectangle.Length * rectangle.Width
}

func calculatePerimeter(rectangle Rectangle) int {
	return 2 * (rectangle.Length + rectangle.Width)
}

func printResults(rectangle Rectangle) {
	fmt.Println("\nLength:", rectangle.Length, "Width:", rectangle.Width)
	fmt.Println("Area:", calculateArea(rectangle))
	fmt.Println("Perimeter:", calculatePerimeter(rectangle))
}

func main() {
	rectangle := Rectangle{Length: 5, Width: 10}

	// - Print the results to the terminal
	printResults(rectangle)

	//* After performing the above requirements, double the size
	//* of the existing rectangle and repeat the calculations
	rectangle.Length *= 2
	rectangle.Width *= 2

	// - Print the new results to the terminal
	printResults(rectangle)
}
