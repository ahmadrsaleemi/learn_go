package main

import (
	"fmt"
	"time"
)

func main() {

	var err error

	fmt.Println("Beginning program.")

	go sayHello()

	fmt.Println("After sayHello function")

	go func() {
		err = doWork()
	}()

	go printNumbers()

	go printLetters()

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Do Work completed successfully")
	}

	time.Sleep(2 * time.Second)
}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Experimenting Go routines in Go language!")
}

func printNumbers() {
	for i:= 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(100*time.Millisecond)
	}
}

func printLetters() {
	startingLetter := 99
	for j:= 0; j< 5; j++ {
		startingLetter += j
		fmt.Println(string(startingLetter))
		time.Sleep(200*time.Millisecond)
	}
}

func doWork() error {
	time.Sleep(1 * time.Second)

	return fmt.Errorf("an error occured in do work!")
}