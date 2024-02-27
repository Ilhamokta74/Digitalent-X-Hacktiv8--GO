package main

import (
	"fmt"
	"unicode/utf8"
)

func Looping1() { // text-to-byte
	city := "Jakarta"

	for i := 0; i < len(city); i++ {
		fmt.Printf("Index : %d, Byte : %d \n", i, city[i])
	}
}

func Looping2() { // byte-to-text
	var city []byte = []byte{74, 97, 107, 97, 114, 116, 97}

	fmt.Println((string(city)))
}

func Looping3() { // rune-by-rune
	str1 := "Makan"

	str2 := "mânca"

	fmt.Printf("Str1 byte length => %d \n", len(str1))
	fmt.Printf("Str2 byte length => %d \n", len(str2))
}

func Looping4() { // rune-by-rune
	str1 := "Makan"

	str2 := "mânca"

	fmt.Printf("Str1 character length => %d \n", utf8.RuneCountInString(str1))
	fmt.Printf("Str2 character length => %d \n", utf8.RuneCountInString(str2))
}

func Looping5() {
	str := "mânca"

	for i, s := range str {
		fmt.Printf("Index => %d, String => %s \n", i, string(s))
	}
}

func main() {
	Looping1()
	Looping2()
	Looping3()
	Looping4()
	Looping5()
}
