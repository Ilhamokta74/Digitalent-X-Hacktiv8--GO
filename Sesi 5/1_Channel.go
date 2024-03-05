package main

import (
	"fmt"
	"time"
)

func introduce(student string, c chan string) {
	result := fmt.Sprintf("hai, my name is %s", student)

	c <- result
}

func print(c chan string) {
	fmt.Println(<-c)
}

// ==================================
func ChannelImplement() {
	c := make(chan string)

	go introduce("Aralie", c)

	go introduce("Nanda", c)

	go introduce("Mailo", c)

	msg1 := <-c
	fmt.Println(msg1)

	msg2 := <-c
	fmt.Println(msg2)

	msg3 := <-c
	fmt.Println(msg3)

	close(c)
}

func ChannelAnonymous() {
	c := make(chan string)

	students := []string{"Aralie", "Mailo", "Indah"}

	for _, v := range students {
		go func(student string) {
			fmt.Println("Student ", student)
			result := fmt.Sprintf("Hai, My name is %s", student)
			c <- result
		}(v)
	}

	for i := 1; i <= 3; i++ {
		print(c)
	}

	close(c)
}

func ChannelDirection() {
	c := make(chan string)

	students := []string{"Aralie", "Mailo", "Indah"}

	for _, v := range students {
		go introduce(v, c)
	}

	for i := 1; i <= 3; i++ {
		print(c)
	}

	close(c)
}

func ChannelUnbuffered() {
	c1 := make(chan int)

	go func(c chan int) {
		fmt.Println("Func goroutine starts sending data into the channel")
		c <- 10
		fmt.Println("Func goroutine after sending data into the channel")
	}(c1)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("Main goroutine starts receiving data")
	d := <-c1
	fmt.Println("Main goroutine receved data : ", d)

	close(c1)
	time.Sleep(time.Second)
}

func ChannelBuffered() {
	c1 := make(chan int, 3)

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Func goroutine #%d starts sending data into the channel \n", i)
			c <- i
			fmt.Printf("Func goroutine #%d after sending data into the channel \n", i)
		}

		close(c)

	}(c1)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	for v := range c1 {
		fmt.Println("Main goroutine received calue from channel : ", v)
	}
}

func ChannelSelect() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)

		c1 <- "Hello!"
	}()

	go func() {
		time.Sleep(1 * time.Second)

		c1 <- "Salut!"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received : ", msg1)
		case msg2 := <-c2:
			fmt.Println("Received : ", msg2)
		}
	}
}

func main() {
	// ChannelImplement()
	// ChannelAnonymous()
	// ChannelDirection()
	// ChannelUnbuffered()
	// ChannelBuffered()
	// ChannelSelect()
}
