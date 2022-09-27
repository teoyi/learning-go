package main

import "fmt"

func main() {
	a := (42 == 43)
	b := (42 != 43)
	c := (42 >= 43)
	d := (42 <= 43)
	e := (42 > 3)
	f := (42 < 3)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
