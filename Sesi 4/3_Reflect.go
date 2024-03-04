package main

import (
	"fmt"
	"reflect"
)

func IdentifyData() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("Tipe Variabel : ", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Nilai Variabel : ", reflectValue.Int())
	}
}

func AccessValue() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("Tipe Variabel : ", reflectValue.Type())
	fmt.Println("Nilai Variabel : ", reflectValue.Interface())
}

type student struct {
	Name  string
	Grade int
}

func (s *student) SetName(name string) {
	s.Name = name
}

func IdentifyMethod() {
	var s1 = &student{Name: "John Wick", Grade: 2}
	fmt.Println("Nama : ", s1.Name)

	var reflectValue = reflect.ValueOf(s1)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})

	fmt.Println("Nama : ", s1.Name)
}

func main() {
	// IdentifyData()
	// AccessValue()
	// IdentifyMethod()
}
