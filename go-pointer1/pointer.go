package main

import (
	"fmt"
)

func test(x int) {
	fmt.Println(x)
	fmt.Println(&x) // here x memory location is different from main ie is in go everythin is passed by value
	x = 0
}

func zero(x *int) {
	fmt.Printf("Memory address %p \n", x)
	*x = 0 // will make x zero
}

func main() {
	a := 30

	b := &a // here var b is of type integer pointer, which hold the value memory location of integer variable a

	fmt.Println(b)  // called referencing
	fmt.Println(*b) // called de referencing

	*b = 45 // replace the meomory locations value by 45 which refered by a
	fmt.Println(a)

	x := 45
	test(45)
	fmt.Println(&x)
	fmt.Println(x)

	p := 50
	fmt.Println(p)
	zero(&p)
	fmt.Println(p)
}
