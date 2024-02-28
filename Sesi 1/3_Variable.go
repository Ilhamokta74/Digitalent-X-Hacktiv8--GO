package main

import "fmt"

func WithDataType() {
	var name string = "Hacktiv8"
	var age int = 20
	fmt.Println("Nama Saya ==>", name)
	fmt.Println("Umur Saya ==>", age)

	// Update Data
	name = "Ilham Oktavian"
	age = 23
	fmt.Println("Nama Saya ==>", name)
	fmt.Println("Umur Saya ==>", age)
}

func WithoutDataType() {
	var name = "Ilham Oktavian"
	var age = 23

	// Check Type Data
	fmt.Printf(`%T, %T`, name, age)
}

func WithoutDataTypeShort() {
	name := "Ilham Oktavian"
	age := 23

	// Check Type Data
	fmt.Printf(`%T, %T`, name, age)
}

func MultiVariableSamedata() {
	var student1, student2, student3 string = "satu", "Dua", "Tiga"
	var first, second, third int = 1, 2, 3

	fmt.Println(student1, student2, student3)
	fmt.Println(first, second, third)
}

func MultiVariableDifferentdata() {
	var name, age, address = "Aralie", 23, "Jalan Sudirman"
	first, second, third := "1", 2, "3"

	fmt.Println(name, age, address)
	fmt.Println(first, second, third)
}

func UnderscoreVariable() {
	var firstVariable string
	var name, age, address = "Aralie", 23, "Jalan Sudirman"

	_, _, _, _ = firstVariable, name, age, address
}

func UsingPrintf() {
	var name = "Aralie"
	var age = 23
	var address = "Jalan Sudirman"

	fmt.Printf("Halo nama ku %s, Umur Aku Adalah %d, dan Aku Tinggal Di %s", name, age, address)
}

func main() {
	// WithDataType()
	// WithoutDataType()
	// WithoutDataTypeShort()
	// MultiVariableSamedata()
	// MultiVariableDifferentdata()
	// UnderscoreVariable()
	// UsingPrintf()
}
