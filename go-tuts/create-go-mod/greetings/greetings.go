package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) { // note that function name is capitalized here. This meant that it is an exported function that can be called elsewhere
	// if no name was given, return an error message
	if name == "" {
		return "", errors.New("empty name")
	}
	// return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// returns hello that associates to each named person
func Hellos(names []string) (map[string]string, error) {
	// a map of assocaite names with messages
	messages := make(map[string]string)
	// loop through the received slice of names and call hello function on each
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// in the map, assocaite the retrieved message with the name.
		messages[name] = message
	}
	return messages, nil
}

// init sets initial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned message is selected at random
func randomFormat() string {
	// a slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// return a randomly selected message format by specifying a random indexd for the slice of formats
	return formats[rand.Intn(len(formats))]
}
