package main

import (
	"fmt"
)

const p string = "Hello from const outside"

func main() {
	const Q int = 42
	fmt.Println(Q)
	fmt.Println(p)
}
