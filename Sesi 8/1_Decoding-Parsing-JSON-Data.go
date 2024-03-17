package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func Decoding_JSON_To_Struct() {
	var jsonString = `
	{
		"full_name": "Aralie Rachel",
		"email": "aralie@gmail.com",
		"age": 23
		}`

	var result Employee

	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("full_name : ", result.FullName)
	fmt.Println("email : ", result.Email)
	fmt.Println("age : ", result.Age)
}

func Decoding_JSON_To_Map() {
	var jsonString = `
	{
		"full_name": "Aralie Rachel",
		"email": "aralie@gmail.com",
		"age": 23
		}`

	var result map[string]interface{}

	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("full_name : ", result["full_name"])
	fmt.Println("email : ", result["email"])
	fmt.Println("age : ", result["age"])
}

func Decoding_JSON_To_Empty_Interface() {
	var jsonString = `
	{
		"full_name": "Aralie Rachel",
		"email": "aralie@gmail.com",
		"age": 23
		}`

	var temp interface{}

	var err = json.Unmarshal([]byte(jsonString), &temp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result = temp.(map[string]interface{})

	fmt.Println("full_name : ", result["full_name"])
	fmt.Println("email : ", result["email"])
	fmt.Println("age : ", result["age"])
}

func Decoding_JSON_Array_To_Slice_Of_Struct() {
	var jsonString = `[
			{
				"full_name": "Aralie Rachel",
				"email": "aralie@gmail.com",
				"age": 23
			},
			{
				"full_name": "Ananda RHP",
				"email": "ananda@gmail.com",
				"age": 23
			}
		]
	`

	var result []Employee

	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, v := range result {
		fmt.Printf("Index %d : %+v \n", i+1, v)
	}
}

func main() {
	// Decoding_JSON_To_Struct()
	// Decoding_JSON_To_Map()
	// Decoding_JSON_To_Empty_Interface()
	// Decoding_JSON_Array_To_Slice_Of_Struct()
}
