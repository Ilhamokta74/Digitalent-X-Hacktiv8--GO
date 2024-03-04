package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	division string
	person   Person
}

func StructEmbedded() {
	var employee1 = Employee{}

	employee1.person.name = "Aralie"
	employee1.person.age = 23
	employee1.division = "Curriculum Developer"

	fmt.Printf("%+v", employee1)
}

func StructAnonymous() {
	// Anonymous struct tanpa pengisian filed
	var employee1 = struct {
		person   Person
		division string
	}{}
	employee1.person = Person{name: "Aralie", age: 23}
	employee1.division = "Curriculum Developer"

	// Anonymous struct dengan pengisian filed
	var employee2 = struct {
		person   Person
		division string
	}{
		person:   Person{name: "Ananda", age: 23},
		division: "Finance",
	}

	fmt.Printf("Employee1 : %+v \n", employee1)
	fmt.Printf("Employee2 : %+v \n", employee2)
}

func StructSlice() {
	var people = []Person{
		{name: "Aralie", age: 23},
		{name: "Ananda", age: 23},
		{name: "mailo", age: 23},
	}

	for _, v := range people {
		fmt.Printf("%+v \n", v)
	}
}

func StructSliceAnonymous() {
	var employee = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "Aralie", age: 23}, division: "Curriculum Developer"},
		{person: Person{name: "Ananda", age: 23}, division: "Finance"},
		{person: Person{name: "Mailo", age: 21}, division: "Marketing"},
	}

	for _, v := range employee {
		fmt.Printf("%+v \n", v)
	}
}

func main() {
	// StructEmbedded()
	// StructAnonymous()
	// StructSlice()
	// StructSliceAnonymous()
}
