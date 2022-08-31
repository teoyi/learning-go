package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("type of x: %T\n", x)
	fmt.Printf("type of y: %T\n", x)
	fmt.Printf("type of z: %T", x)
	// if the values are empty it is called zero values
}
