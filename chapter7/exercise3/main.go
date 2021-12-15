package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (r Rectangle) perimeter() float64 {
	return 2*r.height + 2*r.width
}

func main() {
	circle := Circle{radius: 3}
	rectangle := Rectangle{width: 5, height: 10.1}
	fmt.Println(">> Circle <<")
	printShape(circle)
	fmt.Println(">> Rectangle <<")
	printShape(rectangle)
}

func printShape(shape Shape) {
	fmt.Printf("Area = %.2f\n", shape.area())
	fmt.Printf("Perimeter = %.2f\n", shape.perimeter())
}
