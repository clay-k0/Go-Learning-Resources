package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	isBoarded    bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	kasey := Passenger{"Kasey", 1, false}
	fmt.Println(kasey)

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 2}
		john = Passenger{Name: "John", TicketNumber: 3}
	)
	fmt.Println(bill, john)

	var jack Passenger
	jack.Name = "Jack"
	jack.TicketNumber = 4

	fmt.Println(jack)

	kasey.isBoarded = true
	bill.isBoarded = true

	if bill.isBoarded {
		fmt.Println("Bill is now on the bus!")
	}

	if kasey.isBoarded {
		fmt.Println("Kasey is now on the bus!")
	}

	jack.isBoarded = true
	busState := Bus{jack}
	fmt.Println(busState)
	fmt.Println(busState.FrontSeat.Name, "sat in the front seat!")
}
