package main

import (
	"fmt"
)

func Map() {
	var person map[string]string // Deklarasi

	person = map[string]string{} // Inisiasi

	person["name"] = "Aralie"

	person["age"] = "23"

	person["address"] = "Jalan Sudirman"

	fmt.Println("Name : ", person["name"])
	fmt.Println("Age : ", person["age"])
	fmt.Println("Address : ", person["address"])
}

func Map2() {
	var person = map[string]string{
		"name":    "Aralie",
		"age":     "23",
		"address": "Jalan Sudirman",
	}

	fmt.Println("Name : ", person["name"])
	fmt.Println("Age : ", person["age"])
	fmt.Println("Address : ", person["address"])
}

func MapLoop() {
	var person = map[string]string{
		"name":    "Aralie",
		"age":     "23",
		"address": "Jalan Sudirman",
	}

	for key, value := range person {
		fmt.Println(key, " : ", value)
	}
}

func MapDelete() {
	var person = map[string]string{
		"name":    "Aralie",
		"age":     "23",
		"address": "Jalan Sudirman",
	}

	fmt.Println("Before Deleting : ", person)

	delete(person, "age")

	fmt.Println("After Deleting : ", person)
}

func MapDetecting() {
	var person = map[string]string{
		"name":    "Aralie",
		"age":     "23",
		"address": "Jalan Sudirman",
	}

	value, exist := person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value Does'nt Exist")
	}

	delete(person, "age")

	value, exist = person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value Has Been Deleted")
	}
}

func MapCombining() {
	var people = []map[string]string{
		{"name": "Aralie", "age": "23"},
		{"name": "Nanda", "age": "23"},
		{"name": "Mailo", "age": "15"},
	}

	for i, person := range people {
		fmt.Printf("Index : %d, Name : %s, age : %s \n", i, person["name"], person["age"])
	}
}

func main() {
	// Map()
	// Map2()
	// MapLoop()
	// MapDelete()
	// MapDetecting()
	// MapCombining()
}
