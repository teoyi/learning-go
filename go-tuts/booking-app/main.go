package main // all code needs to belong to a package.
// first statement of a go file must be package and the name of the package -> smth like declare namespace?

import (
	"booking-app/helper"
	"fmt" // just like python libraries needs to be imported fmt - format
	"sync"
	"time"
)

// here is package level so can be assesed anywhere
var conferenceName = "Go Conference" // to create variables use var prefix to tell go that this is a variable
// a different way to create a var we can use := to replace var, but in this case we cant fix its type
const conferenceTickets = 50 // for constants in the code, we add prefix const
var remainingTickets uint = 50

// instead of a list of string we want a map
// var bookings = make([]map[string]string, 0) // to create an empty slice of maps, we need to indicate the starting size. It is dynamic so putting 0 works

var bookings = make([]UserData, 0) // using struct for our bookings
// TO DEAL WITH MULTI TYPES  we use struct
type UserData struct { // this creates a custom type named UserData
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() { // this is where it tells the compiler the code starts here only 1 main func in go file

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceaName is %T\n", conferenceName, conferenceTickets, remainingTickets) // to voew ots ty[e ]
	// fmt.Print("Hello World") // Print on its own does not exist, thats why fmt needs to be imported for it to work
	// fmt.Println("Hello World with line below")
	greetUsers()

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

	// a data type that lets us store key values is Maps

	//adding new elements

	// for loop time
	for { // indefinite loop
		// for remainingTickets > 0 { // we can add a condition to make it work in indefinite loop

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTickets {
			bookTickets(remainingTickets, userTickets, firstName, lastName, email)
			// to tackle the iossue of main thread not waiting for other threads we use waitgroup
			wg.Add(1)                                               // number here depends on number of threads to be added to the wait group. If more than 1 go exists, we +1 the nuumber
			go sendTickets(userTickets, firstName, lastName, email) // go starts a go routine and introduce concurrency however main thread does not wait for other threads to finish their work to end the aplpication
			if remainingTickets == 0 {
				// end the program here
				fmt.Println("Our conference is booked out. Come back next year.")
				break // breaks the indefinite loop
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short ")
			}
			if !isValidEmail {
				fmt.Println("Email entered is not valid")
			}
			if !isValidTickets {
				fmt.Println("Number of ticekts you entered is invalid")
			}
			// fmt.Printf("Your input data is invalid pls try again.\n", remainingTickets, userTickets)
			// continue // instead of breaking out, skip to the next iteration, not needed if in else block
		}
	}
	wg.Wait() // this tells main thread to wait for threads in wait group to be done so add done command at end of logic
}

func greetUsers() { // need to tell function what type the variable is gonna be
	// fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("Welcome to %v booking application", conferenceName) // space gets added automatically
	fmt.Printf("We have total of %v tickets and %v are still available", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// to execute, use { go run <file name>}
func printFirstNames() []string { // need to define type of returned values
	firstNames := []string{}
	for _, booking := range bookings { // range is needed to iterate over elements, _ is used to replace unused variables, index is in its place range gives index and value
		// var names = strings.Fields(booking) // split strings with white space
		// var firstName = names[0]
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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
	return firstName, lastName, email, userTickets
}

func bookTickets(remainingTickets uint, userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName // for arary

	// create a map for a user
	// var userData = make(map[string]string) //  craete empty map and their types. Inside brackets is key outside is value. Make is built in to tell it to create a map cannot mix data types in map
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// to add value into map
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) // formatUint takes in a value and formats it into a decimal number, second argument is base 10

	// bookings = append(bookings, firstName+" "+lastName)
	bookings = append(bookings, userData)

	// fmt.Printf("The whole array: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("Array Type: %T\n", bookings)
	// fmt.Printf("Array length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n ", firstName, lastName, userTickets, email)
	fmt.Printf("%v remaining from %v\n", remainingTickets, conferenceName)

	firstNames := printFirstNames()
	fmt.Printf("These are the list of bookers %v\n", firstNames)
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)                                                       // to simulate that this function takes time to run
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName) //sprintf allows asving to var
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("##########################################")
	wg.Done()
}
