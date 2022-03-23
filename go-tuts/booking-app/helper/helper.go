// package main

package helper

// best define functions here on the package level
// when using go run, we need to put helper.go as well to run both at the same time

// alternatively, we can use go run using . or a folder that contains all the code
import "strings"

// note to export, make the first letter capitalized
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets < remainingTickets

	return isValidName, isValidEmail, isValidTickets
}
