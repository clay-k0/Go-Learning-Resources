package main

import "fmt"

type Room struct {
	name      string
	isCleaned bool
}

func checkIfClean(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.isCleaned {
			fmt.Println("Room is cleaned")
		} else {
			fmt.Println("Room is dirty")
		}
	}
}

func main() {
	rooms := [...]Room{
		{name: "Office"},
		{name: "Warehouse"},
		{name: "Reception"},
		{name: "Ops"},
	}

	checkIfClean(rooms)

	fmt.Println("Cleaning...")

	rooms[0].isCleaned = true
	rooms[1].isCleaned = true

	checkIfClean(rooms)
}
