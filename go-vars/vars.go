package main

import (
	"fmt"
)

func main() {
	a := 10
	b := "String"
	c := 2.4
	d := true

	fmt.Printf("%T \t %v \n", a, a)
	fmt.Printf("%T \t %v \n", b, b)
	fmt.Printf("%T \t %v \n", c, c)
	fmt.Printf("%T \t %v \n", d, d)
}
