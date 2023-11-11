package main

import "fmt"

func main() {
	var myName = "Clayton"
	fmt.Println("my name is", myName, myName)

	var otherName string = "Mary"
	fmt.Println("Hello,", otherName)

	username := "admin"
	fmt.Println("username:", username)

	var sum int
	fmt.Println("sum is currently", sum)

	part1, other := 1, 5
	fmt.Println("part1 is", part1, "other is", other)

	part2, other := 2, 0
	fmt.Println("part2 is", part2, "other is", other)

	sum = part1 + part2
	fmt.Println("sum is now", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("Lesson Name:", lessonName)
	fmt.Println("Lesson Type:", lessonType)

	word1, word2, _ := "Hello", "World", "!"
	fmt.Println(word1, word2)
}
