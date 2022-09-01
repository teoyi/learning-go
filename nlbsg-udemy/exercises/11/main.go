package main

import "fmt"

const (
	a = 2022 - iota
	b = 2022 - iota
	c = 2022 - iota
	d = 2022 - iota
)

func main() {
	fmt.Println(a, b, c, d)
}
