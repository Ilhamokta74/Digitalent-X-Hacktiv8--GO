package main

import (
	"fmt"
	"os"
)

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("Defer func starts to execute")
}

// ============================================================
func Defer1() {
	defer fmt.Println("defer function starts to execute")
	fmt.Println("Hai everyone")
	fmt.Println("Welcome back to Go learning center")
}

func Defer2() {
	callDeferFunc()
	fmt.Println("Hai everyone!!")
}

func Exit() {
	defer fmt.Println("Invoke with defer")
	fmt.Println("Before Exiting")
	os.Exit(1)
}

func main() {
	// Defer1()
	// Defer2()
	// Exit()
}
