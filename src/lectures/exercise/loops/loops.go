//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func main() {

	for i := 1; i <= 50; i++ {
		divisibleByThree := i%3 == 0
		divisibleByFive := i%5 == 0

		if divisibleByFive && divisibleByThree {
			//  - Print "Fizz" if the integer is divisible by 3
			fmt.Println("FizzBuzz")
		} else if divisibleByThree {
			//  - Print "Buzz" if the integer is divisible by 5
			fmt.Println("Fizz")
		} else if divisibleByFive {
			//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
