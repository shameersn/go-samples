package main

import (
	"fmt"
)

func loop() {
	sum := 0
	for i := 0; i < 19; i++ {
		sum += i
	}
	fmt.Println(sum)
}
func nested() {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(i, " - ", j)
		}
	}
}

func shortFor() {
	i := 0
	for i < 5 {
		fmt.Println("*")
		i++
	}
}

func breakFor() {
	i := 0
	for {
		fmt.Println("$")
		if i > 5 {
			break // break stop loop and out of the loop,  also continue will skip the current iteration
		}
		i++
	}
}
func main() {
	loop()
	nested()
	shortFor()
	breakFor()
}
