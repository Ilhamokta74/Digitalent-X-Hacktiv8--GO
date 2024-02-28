package main

import (
	"fmt"
	"sync"
)

func main() {
	fruits := []string{"Apple", "Manggo", "Durian", "Rambutan"}

	var wg sync.WaitGroup

	for index, fruit := range fruits {
		wg.Add(1)
		go printFruit(index, fruit, &wg)
	}

	wg.Wait()
}

func printFruit(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("Index => %d, fruit => %s \n", index, fruit)
	wg.Done()
}
