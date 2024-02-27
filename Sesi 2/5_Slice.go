package main

import (
	"fmt"
	"strings"
)

func Slice() {
	var fruits = []string{"Apple", "Banana", "Mango"}

	_ = fruits
}

func SliceMake() {
	var fruits = make([]string, 3)

	_ = fruits

	fmt.Printf("%#v", fruits)
}

func SliceAppend1() {
	var fruits = make([]string, 4)

	_ = fruits

	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Mango"
	fruits[3] = "Durian"

	fmt.Printf("%#v", fruits)
}

func SliceAppend2() {
	var fruits = make([]string, 3)

	_ = fruits

	fruits = append(fruits, "Apple", "Banana", "Mango")

	fmt.Printf("%#v", fruits)
}

func SliceAppend3() {
	var fruits1 = []string{"Apple", "Banana", "Mango"}
	var fruits2 = []string{"Durian", "Pineapple", "Starfruit"}

	var fruits = append(fruits1, fruits2...)

	fmt.Printf("%#v", fruits)
}

func SliceCopy() {
	var fruits1 = []string{"Apple", "Banana", "Mango"}
	var fruits2 = []string{"Durian", "Pineapple", "Starfruit"}

	nn := copy(fruits1, fruits2)

	fmt.Println("Fruits1 => ", fruits1)
	fmt.Println("Fruits2 => ", fruits2)
	fmt.Println("Copied Element", nn)
}

func SliceWithIndex() {
	var fruits1 = []string{"Apple", "Banana", "Mango", "Durian", "Pineapple", "Starfruit"}

	var fruits2 = fruits1[1:4]
	fmt.Printf("%#v \n", fruits2)

	var fruits3 = fruits1[0:]
	fmt.Printf("%#v \n", fruits3)

	var fruits4 = fruits1[:3]
	fmt.Printf("%#v \n", fruits4)

	var fruits5 = fruits1[:]
	fmt.Printf("%#v \n", fruits5)
}

func SliceCombineAndAppend() {
	var fruits1 = []string{"Apple", "Banana", "Mango", "Durian", "Pineapple", "Starfruit"}

	fruits1 = append(fruits1[:3], "Rambutan")

	fmt.Printf("%#v", fruits1)
}

func SliceBacking() {
	var fruits1 = []string{"Apple", "Banana", "Mango", "Durian", "Pineapple", "Starfruit"}

	var fruits2 = fruits1[2:4]

	fruits2[0] = "Rambutan"

	fmt.Println("Fruits1 => ", fruits1)
	fmt.Println("Fruits2 => ", fruits2)
}

func SliceCap() {
	var fruits1 = []string{"Apple", "Mango", "Durian", "Banana"}

	fmt.Println("Fruits1 Cap : ", cap(fruits1))
	fmt.Println("Fruits1 Len : ", len(fruits1))
	fmt.Println(strings.Repeat("#", 20))

	var fruits2 = fruits1[0:3]

	fmt.Println("Fruits2 Cap : ", cap(fruits2))
	fmt.Println("Fruits2 Len : ", len(fruits2))
	fmt.Println(strings.Repeat("#", 20))

	var fruits3 = fruits1[1:]

	fmt.Println("Fruits3 Cap : ", cap(fruits3))
	fmt.Println("Fruits3 Len : ", len(fruits3))
}

func SliceNewBacking() {
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println("Cars : ", cars)
	fmt.Println("New Cars : ", newCars)
}

func main() {
	Slice()
	SliceMake()
	SliceAppend1()
	SliceAppend2()
	SliceAppend3()
	SliceCopy()
	SliceWithIndex()
	SliceCombineAndAppend()
	SliceBacking()
	SliceCap()
	SliceNewBacking()
}
