package main

import (
	"fmt"
	"strings"
)

func ClosureGenap() {
	var evenNumbers = func(numbers ...int) []int {
		var result []int

		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}

		return result
	}

	var numbers = []int{4, 93, 77, 10, 52, 22, 34}

	fmt.Println(evenNumbers(numbers...))
}

func ClosureIfElse() {
	var isPolindrome = func(str string) bool {
		var temp string = ""

		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}
		return temp == str
	}("katak")

	fmt.Println(isPolindrome)
}

func ClosureReturnValue(students []string) func(string) string {
	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = i
				break
			}
		}
		if student == "" {
			return fmt.Sprintf("%s does'ntexist!!!", s)
		}
		return fmt.Sprintf("We found %s at position %d", s, position+1)
	}
}

type isOddNum func(int) bool

func ClosureGanjil(numbers []int, callback isOddNum) int {
	var totalOddNumbers int

	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}

	return totalOddNumbers
}

func main() {
	// ClosureGenap()
	// ClosureIfElse()

	// var studentLists = []string{"Aralie", "Nanda", "Mailo", "Shannel", "Marco"}
	// var find = ClosureReturnValue(studentLists)
	// fmt.Println(find("Aralie"))

	// var numbers = []int{2, 5, 8, 10, 3, 99, 23}
	// var find = ClosureGanjil(numbers, func(number int) bool {
	// 	return number%2 != 0
	// })
	// fmt.Println("Total odd numbers : ", find)
}
