//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Perform the following:
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

// * Create a function to print out the contents of the assembly line
func printParts(slice []Part) {
	fmt.Println()
	fmt.Println("---Assembly Line---")
	for i := 0; i < len(slice); i++ {
		part := slice[i]
		fmt.Println(">", part)
	}
}

func main() {
	//  - Create an assembly line having any three parts
	assemblyLine := []Part{"Conveyor Belts", "Workers", "Carts"}
	printParts(assemblyLine)

	//  - Add two new parts to the line
	assemblyLine = append(assemblyLine, "Bins", "Safety Equipment")
	fmt.Println()
	fmt.Println("Added", assemblyLine[3], "+", assemblyLine[4])
	printParts(assemblyLine)

	//  - Slice the assembly line so it contains only the two new parts
	assemblyLine = assemblyLine[3:]
	fmt.Println()
	fmt.Println("Removed All But New Parts")
	printParts(assemblyLine)
}
