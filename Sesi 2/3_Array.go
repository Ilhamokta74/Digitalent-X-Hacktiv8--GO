package main

import (
	"fmt"
	"strings"
)

func Array() {
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var strings = [3]string{"Aralie", "Nanda", "Mailo"}

	fmt.Printf("%#v \n", numbers)
	fmt.Printf("%#v \n", strings)
}

func ArrayModify() {
	var fruits = [3]string{"Apel", "Pisang", "Mangga"}
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Mango"

	fmt.Printf("%#v \n", fruits)
}

func ArrayLoop() {
	var fruits = [3]string{"Apple", "Banana", "Mango"}

	// Cara Pertama
	for i, v := range fruits {
		fmt.Printf("Index : %d, Value : %s \n", i, v)
	}

	fmt.Println(strings.Repeat("#", 25))

	// Cara Kedua
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Index : %d, Value : %s \n", i, fruits[i])
	}
}

func ArrayMultidimensional() {
	balance := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balance {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}

func main() {
	// Array()
	// ArrayModify()
	// ArrayLoop()
	// ArrayMultidimensional()
}
