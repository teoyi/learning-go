package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// set properties of the predefined logger, including the log entry prefix and a flag to disable printing the time, source file, and line number .
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Ayano", "Luke", "Honey"}
	// Request a greeting message
	// get a greeting message and print it.
	messages, err := greetings.Hellos(names)
	// If an error was returned, print it to the console and exist the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message to the console.
	fmt.Println(messages)
}
