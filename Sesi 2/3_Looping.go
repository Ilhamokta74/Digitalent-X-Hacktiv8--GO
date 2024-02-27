package main

import "fmt"

func For() {
	for i := 0; i < 3; i++ {
		fmt.Println("Angka", i)
	}
}

func ForContinueAndBreak() {
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}
}

func NestedLoop() {
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

func LoopingLabel() {
	for i := 0; i < 3; i++ {
		fmt.Println("Perulangan ke - ", i+1)
		for j := i; j < 3; j++ {
			if i == 2 {
				break
			}
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

func main() {
	For()
	ForContinueAndBreak()
	NestedLoop()
	LoopingLabel()
}
