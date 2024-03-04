package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	permeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) permeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) permeter() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func print(t string, s shape) {
	fmt.Printf("%s Area : %v \n", t, s.area())
	fmt.Printf("%s Perimeter %v \n", t, s.permeter())
}

func Interface1() {
	var c1 shape = circle{radius: 5}
	var r1 shape = rectangle{width: 3, height: 2}

	fmt.Printf("Type of c1 : %T \n", c1)
	fmt.Printf("Type of r1 : %T", r1)
}

func Interface2() {
	var c1 shape = circle{radius: 5}
	var r1 shape = rectangle{width: 3, height: 2}

	fmt.Println("Circle Area : ", c1.area())
	fmt.Println("Circle Perimeter ", c1.permeter())

	fmt.Println("Rectangle Area : ", r1.area())
	fmt.Println("Rectangle Perimeter ", r1.permeter())
}

func Interface3() {
	var c1 shape = circle{radius: 5}
	var r1 shape = rectangle{width: 3, height: 2}

	print("Rectangle", c1)
	print("Circle", r1)
}

func InterfaceTypeAssertion() {
	var c1 shape = circle{radius: 5}

	c1.(circle).volume()

	value, ok := c1.(circle)

	if ok == true {
		fmt.Printf("Circle Value : %+v \n", value)
		fmt.Printf("Circle Volume : %f \n", value.volume())
	}
}

func main() {
	// Interface1()
	// Interface2()
	// Interface3()
	// InterfaceTypeAssertion()
}
