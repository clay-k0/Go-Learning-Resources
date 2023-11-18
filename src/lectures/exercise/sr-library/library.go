//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Title string
type Name string

type LendAudit struct {
	checkInTime  time.Time
	checkOutTime time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int
	lended int
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkInTime.IsZero() {
			returnTime = "[not yet returned]"
		} else {
			returnTime = audit.checkInTime.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkOutTime.String(), "through", returnTime)
	}
}

func printMemberAudits(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

func createBook(library *Library, title Title, total int, lended int) {
	library.books[title] = BookEntry{
		total,
		lended,
	}
}

func printBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "| total:", book.total, "| lended:", book.lended)
	}
	fmt.Println()
}

func checkOutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("book not found")
		return false
	}
	if book.lended == book.total {
		fmt.Println("no more books available to lend")
		return false
	}
	book.lended += 1
	library.books[title] = book

	member.books[title] = LendAudit{checkOutTime: time.Now()}
	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("book not found")
		return false
	}

	audit, found := member.books[title]
	if !found {
		fmt.Println("member did not check this book out")
		return false
	}

	book.lended -= 1
	library.books[title] = book

	audit.checkInTime = time.Now()
	member.books[title] = audit
	return true
}

func main() {
	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	createBook(&library, "Learning Go", 3, 0)
	createBook(&library, "Learning Rust", 2, 0)
	createBook(&library, "Learning Python", 1, 0)
	createBook(&library, "Learning JavaScript", 5, 0)

	library.members["Clay"] = Member{"Clay", make(map[Title]LendAudit)}
	library.members["Dan"] = Member{"Dan", make(map[Title]LendAudit)}
	library.members["Tom"] = Member{"Tom", make(map[Title]LendAudit)}

	fmt.Println("\ninitial:")
	printBooks(&library)
	printMemberAudits(&library)

	member := library.members["Clay"]
	checkedOut := checkOutBook(&library, "Learning Go", &member)
	fmt.Println("\ncheck out a book:")
	if checkedOut {
		printBooks(&library)
		printMemberAudits(&library)
	}

	returned := returnBook(&library, "Learning Go", &member)
	fmt.Println("\ncheck in a book:")
	if returned {
		printBooks(&library)
		printMemberAudits(&library)
	}

}
