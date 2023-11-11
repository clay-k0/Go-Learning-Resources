package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a + b + c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 53, 71, 100

	if quiz1 > quiz2 {
		fmt.Println("quiz1 (", (quiz1), ") is higher than quiz2 (", (quiz2), ")")
	} else if quiz1 < quiz2 {
		fmt.Println("quiz2 (", (quiz2), ") is higher than quiz1 (", (quiz1), ")")
	} else {
		fmt.Println("quiz1 (", (quiz1), ") and quiz2 (", (quiz2), ") are equal")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("acceptable grades")
	} else {
		fmt.Println("do better")
	}

}
