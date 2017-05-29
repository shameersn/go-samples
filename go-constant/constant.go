package main

import (
	"fmt"
)

const p string = "Hello from const outside"

const (
	run = "Go to your home"
	age = 23
)

// const using iota
const (
	A = iota
	B
	C
)

func main() {
	const Q int = 42
	fmt.Println(Q)
	fmt.Println(p)
	fmt.Println(run)
	fmt.Println(age)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
