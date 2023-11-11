package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {

	switch p := price(); {
	case p < 2:
		fmt.Println("cheap item")
	case p < 10:
		fmt.Println("moderately priced item")
	default:
		fmt.Println("expensive item")
	}

	ticket := Economy

	switch ticket {
	case Economy:
		fmt.Println("economy seating")
	case Business:
		fmt.Println("business seating")
	case FirstClass:
		fmt.Println("first class seating")
	default:
		fmt.Println("other seating")
	}

}
