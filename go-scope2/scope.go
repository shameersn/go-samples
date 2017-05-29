package main

import (
	"fmt"
)

func wraper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wraper()

	fmt.Println(increment())
	fmt.Println(increment())
}
