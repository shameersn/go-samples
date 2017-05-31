package main

import (
	"fmt"
)

var x = 12

func main() {
	a := 10
	b := "String"
	c := 2.4
	d := true
	var e = "string with type asigned by go"

	fmt.Printf("%T \t %v \n", a, a)
	fmt.Printf("%T \t %v \n", b, b)
	fmt.Printf("%T \t %v \n", c, c)
	fmt.Printf("%T \t %v \n", d, d)
	fmt.Printf("%T \t %v \n", x, x)
	fmt.Printf("%T \t %v \n", e, e)
}
