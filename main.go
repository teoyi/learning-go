package main // all code needs to belong to a package.
// first statement of a go file must be package and the name of the package -> smth like declare namespace?

import (
	"fmt" // just like python libraries needs to be imported fmt - format
	"strings"
)

func main() { // this is where it tells the compiler the code starts here only 1 main func in go file
	var conferenceName = "Go Conference" // to create variables use var prefix to tell go that this is a variable
	// a different way to create a var we can use := to replace var, but in this case we cant fix its type
	const conferenceTickets = 50 // for constants in the code, we add prefix const
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceaName is %T\n", conferenceName, conferenceTickets, remainingTickets) // to voew ots ty[e ]
	// fmt.Print("Hello World") // Print on its own does not exist, thats why fmt needs to be imported for it to work
	// fmt.Println("Hello World with line below")
	fmt.Println("Welcome to", conferenceName, "booking application") // space gets added automatically
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	// instead of typing it out like above, we can use template string
	fmt.Println("")
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	// note however Printf does not have new line so we will have to include \n

	// Array for bookings
	// var bookings = [50]string{} // value inside states how many elements it can hold followed by data type
	// note that elements are used with curly braces and only same data type can be stored

	// in the case of empty array
	// var bookings [50]string
	// to create a slice, its a dynamic array like list in python
	// can be done with
	var bookings []string // whreby you dont define  a size
	//adding new elements

	// for loop time
	for { // indefinite loop 
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name
		// in go, it requires hard typing
		// two basic data types
		// String for textual
		// Integers for whole numbers but there are different types of numeric data types
		// If we assign a value when variable is created, go can infer
		// else we need to tell explicitly typed

		// pointer points to memory address
		// fmt.Println(&remainingTickets) // this gives the memory address

		fmt.Println("Enter your first name")
		fmt.Scan(&firstName) // get value from user input by waiting for input
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName) // get value from user input by waiting for input
		fmt.Println("Enter your email")
		fmt.Scan(&email) // get value from user input by waiting for input
		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets) // get value from user input by waiting for input

		remainingTickets = remainingTickets - userTickets
		// bookings[0] = firstName + " " + lastName // for arary
		bookings = append(bookings, firstName+" "+lastName)

		// fmt.Printf("The whole array: %v\n", bookings)
		// fmt.Printf("The first value: %v\n", bookings[0])
		// fmt.Printf("Array Type: %T\n", bookings)
		// fmt.Printf("Array length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n ", firstName, lastName, userTickets, email)
		fmt.Printf("%v remaining from %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for index, booking := range bookings {// range is needed to iterate over elements 
			var strings.Fields(booking) // split strings with white space
			var firstName = names[0] 
		}
			fmt.Printf("These are the list of bookers %v\n", bookings)
	}

}

// to execute, use { go run <file name>}
