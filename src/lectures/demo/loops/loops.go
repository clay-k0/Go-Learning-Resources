package main

import "fmt"

func main() {
	sum := 0
	fmt.Println("sum is currently", sum)

	for i := 1; i <= 10; i++ {
		sum += i
		fmt.Println("sum is now", sum)
	}

	for sum > 10 {
		sum -= 5
		fmt.Println("(decrement) sum is now", sum)
	}
}
