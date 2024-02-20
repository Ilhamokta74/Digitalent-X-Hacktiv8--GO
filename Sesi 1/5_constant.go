package main

import "fmt"

func Constant() {
	const fullname string = "Aralie Rachel"

	fmt.Printf(fullname)
}

func ArithmeticOperators() {
	var value = (2 + 2) * 3

	fmt.Println(value)
}

func RelationalOperators() {
	var firstCondition bool = 2 < 3
	var secondCondition bool = "joey" == "Joey"
	var thirdCondition bool = 10 != 2.3
	var fourCondition bool = 11 <= 11

	fmt.Println("First Condition : ", firstCondition)
	fmt.Println("Second Condition : ", secondCondition)
	fmt.Println("Third Condition : ", thirdCondition)
	fmt.Println("Four Condition : ", fourCondition)
}

func LogicalOperators() {
	var right = true
	var wrong = false

	var wrongAndRight = wrong && right
	fmt.Printf("Wrong && Right \t(%t) \n", wrongAndRight)

	var wrongOrRight = wrong || right
	fmt.Printf("Wrong || Right \t(%t) \n", wrongOrRight)

	var wrongReverse = !wrong
	fmt.Printf("Reverse Wrong \t(%t) \n", wrongReverse)
}

func main() {
	Constant()
	ArithmeticOperators()
	RelationalOperators()
	LogicalOperators()
}
