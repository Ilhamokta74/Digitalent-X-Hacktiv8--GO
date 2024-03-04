package main

import (
	"fmt"
	"strings"
)

func Pointer1() { // Memory Address
	var firstNumber int = 4
	var secondNumber *int = &firstNumber

	fmt.Println("firstNumber (Value) : ", firstNumber)
	fmt.Println("firstNumber (Memory Address) : ", &firstNumber)

	fmt.Println("secondNumber (Value) : ", *secondNumber)
	fmt.Println("secondNumber (Memory Address) : ", secondNumber)
}

func Pointer2() { // Changing value through pointer
	var firstPerson string = "John"
	var secondPerson *string = &firstPerson

	fmt.Println("FirstPerson (Value) : ", firstPerson)
	fmt.Println("FirstPerson (Memory Address) : ", &firstPerson)
	fmt.Println("SecondPerson (Value) : ", *secondPerson)
	fmt.Println("SecondPerson (Memory Address) : ", secondPerson)

	fmt.Println("\n", strings.Repeat("#", 60), "\n")

	*secondPerson = "Doe"

	fmt.Println("FirstPerson (Value) : ", firstPerson)
	fmt.Println("FirstPerson (Memory Address) : ", &firstPerson)
	fmt.Println("SecondPerson (Value) : ", *secondPerson)
	fmt.Println("SecondPerson (Memory Address) : ", secondPerson)
}

func Pointer3(number *int) {
	*number = 20
}

func main() {
	// Pointer1()
	// Pointer2()

	// var a int = 10
	// fmt.Println("Before : ", a)
	// Pointer3(&a)
	// fmt.Println("After : ", a)
}
