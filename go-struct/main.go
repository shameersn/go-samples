package main

import (
	"fmt"
	"math"
)

func main() {
	dog := Animal{name: "hugo", legs: 4, age: 20}
	fmt.Println(dog.name)
	fmt.Println(dog.sayHi())
	cat := Animal{"tom", 4, 29}
	fmt.Println("Cats Age :", cat.age)
	fmt.Println(cat.sayHi())

	rectangle := Rectangle{10, 10}
	circle := Circle{10}

	fmt.Println("Area of rectangle:", getArea(rectangle))
	fmt.Println("Area of circle:", getArea(circle))
}

// Animal type
type Animal struct {
	name string
	legs int
	age  int
}

func (animal *Animal) sayHi() string {
	return "Hello " + animal.name
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	height, width float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func getArea(s Shape) float64 {
	return s.area()
}
