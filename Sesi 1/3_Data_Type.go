package main

import "fmt"

func IntegerNormal() {
	// Integer Normal
	var first = 89
	var second = -12

	fmt.Printf("Type Data first : %T \n", first)
	fmt.Printf("Type Data second : %T \n", second)
}

func IntegerNextLevel() {
	// Integer Next Level
	var first uint8 = 89
	var second int8 = -12

	fmt.Printf("Type Data first : %T \n", first)
	fmt.Printf("Type Data second : %T \n", second)
}

func Float() {
	var decimalNumber float32 = 3.63

	fmt.Printf("Desimal Number %f \n", decimalNumber)
	fmt.Printf("Desimal Number %.3f \n", decimalNumber)
}

func Boolean() {
	var condition bool = true

	fmt.Printf("is It permitted ? %t \n", condition)
}

func String() {
	var message string = "Halo"

	fmt.Println(message)
}

func main() {
	// IntegerNormal()
	// IntegerNextLevel()
	// Float()
	// Boolean()
	// String()
}
