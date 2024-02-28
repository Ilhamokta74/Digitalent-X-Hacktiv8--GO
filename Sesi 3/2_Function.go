package main

import (
	"fmt"
	"math"
	"strings"
)

func Function(name string, age uint8) {
	fmt.Printf("Hello There, my name is %s and I'm %d years old", name, age)
}

func FunctionReturn(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")

	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result
}

func FunctionReturnMultiple(d float64) (float64, float64) {
	// Menghitung Luas
	var area float64 = math.Pi * math.Pow(d/2, 2)

	// Menghitung Keliling
	var circumference = math.Pi * d

	return area, circumference
}

func FunctionReturnPredefined(d float64) (area float64, circumference float64) {
	// Menghitung Luas
	area = math.Pi * math.Pow(d/2, 2)

	// Menghitung Keliling
	circumference = math.Pi * d

	return
}

func FunctionVariadic1(names ...string) []map[string]string {
	var result []map[string]string
	for i, v := range names {
		key := fmt.Sprintf("Student %d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}

	return result
}

func FunctionVariadic2(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}
	return total
}

func FunctionVariadic3(name string, favfoods ...string) {
	mergeFavFoods := strings.Join(favfoods, ",")

	fmt.Println("Hello There!!! I'm ", name)
	fmt.Println("I Really Love To Eat ", mergeFavFoods)
}

func main() {
	// Function("Aralie", 23)

	// var names = []string{"Aralie", "Jordan"}
	// var printMsg = FunctionReturn("Heii", names)
	// fmt.Println(printMsg)

	// var diameter float64 = 15
	// var area, circumference = FunctionReturnMultiple(diameter)
	// fmt.Println("Area : ", area)
	// fmt.Println("Circumference : ", circumference)

	// var diameter float64 = 15
	// var area, circumference = FunctionReturnPredefined(diameter)
	// fmt.Println("Area : ", area)
	// fmt.Println("Circumference : ", circumference)

	// studentLists := FunctionVariadic1("Aralie", "Nanda", "Mailo", "Schannel", "Marco")
	// fmt.Printf("%v", studentLists)

	// numberLists := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// result := FunctionVariadic2(numberLists...)
	// fmt.Println("Result : ", result)

	// FunctionVariadic3("Aralie", "Pasta", "Ayam Geprek", "Ikan Roa", "Sate Padang")
}
