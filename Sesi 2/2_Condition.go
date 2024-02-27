package main

import "fmt"

func IfElse() {
	var currentYear = 2021

	if age := currentYear - 1998; age < 17 {
		fmt.Println("kamu Belum Boleh Membuat kartu SIM")
	} else {
		fmt.Println("Kamu Boleh Membuat Kartu SIM")
	}
}

func Switch() {
	var score = 8

	switch score {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not Bad")
	}
}

func SwitchWithRelational() {
	var score = 6

	switch {
	case score == 8:
		fmt.Println("Perfect")
	case (score < 8) && (score > 3):
		fmt.Println("Not Bad")
	default:
		fmt.Println("Study Harder")
		fmt.Println("You Need To Learn More")
	}
}

func SwitchWithFallthrough() {
	var score = 6

	switch {
	case score == 8:
		fmt.Println("Perfect")
	case (score < 8) && (score > 3):
		fmt.Println("Not Bad")
		fallthrough
	case score < 5:
		fmt.Println("It Is Ok, But Please Study harder")
	default:
		fmt.Println("Study Harder")
		fmt.Println("You Need To Learn More")
	}
}

func Nested() {
	var score = 10

	if score > 7 {
		switch score {
		case 10:
			fmt.Println("Perfect!")
		default:
			fmt.Println("Nice!")
		}
	} else {
		if score == 5 {
			fmt.Println("Not Bad")
		} else if score == 3 {
			fmt.Println("Keep Trying")
		} else {
			fmt.Println("You can Do It")
			if score == 0 {
				fmt.Println("Try Harder!")
			}
		}
	}
}

func main() {
	IfElse()
	Switch()
	SwitchWithRelational()
	SwitchWithFallthrough()
	Nested()
}
