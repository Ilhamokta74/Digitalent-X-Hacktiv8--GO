package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	name     string
	age      int
	division string
}

func StructGivingValue() {
	var employee Employee

	employee.name = "Aralie"
	employee.age = 23
	employee.division = "Curriculum Developer"

	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)
}

func StructInitializing() {
	var employee1 = Employee{}

	employee1.name = "Aralie"
	employee1.age = 23
	employee1.division = "Curriculum Developer"

	var employee2 = Employee{name: "Ananda", age: 23, division: "Finance"}

	fmt.Printf("Employee1 : %+v \n", employee1)
	fmt.Printf("Employee2 : %+v \n", employee2)
}

func StructPointer() {
	var employee1 = Employee{name: "Aralie", age: 23, division: "Curriculum Developer"}

	var employee2 *Employee = &employee1

	fmt.Println("Employee1 Name : ", employee1.name)
	fmt.Println("Employee2 Name : ", employee2.name)

	fmt.Println(strings.Repeat("#", 40))

	employee2.name = "Ananda"

	fmt.Println("Employee1 Name : ", employee1.name)
	fmt.Println("Employee2 Name : ", employee2.name)
}

func main() {
	// StructGivingValue()
	// StructInitializing()
	// StructPointer()
}
