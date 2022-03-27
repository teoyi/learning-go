package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// set properties of the predefined logger, including the log entry prefix and a flag to disable printing the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// a slice of names
	names := []string{"Ayano", "Honey", "Puppy"}

	// get a greeting message and print it
	message, err := greetings.Hellos(names)
	// if error returne, print and exit
	if err != nil {
		log.Fatal(err)
	}
	// if no error, print the returned message
	fmt.Println(message)
}
