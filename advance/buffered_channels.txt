package main

import (
	"fmt"
	"time"
)

func main() {
	//	BLOCKING ON RECEIVE ONLY IF THE BUFFER IS EMPTY
	ch := make(chan int, 2)

	go func(){
		time.Sleep(2 * time.Second)
		ch <- 1
	}()

	fmt.Println("value: ", <-ch)
	fmt.Println("End of program!")
}

// func main() {
//	BLOCKING EXAMPLE ON SEND ONLY IF THE BUFFER IS FULL
// 	ch := make(chan int, 2)
// 	ch <- 1
// 	ch <- 2

// 	go func() {
// 		for value := range ch {
// 			fmt.Println("Value: ", value)
// 		}
// 	}()
// 	ch <- 3
// 	ch <- 4
// 	ch <- 5
// 	fmt.Println("Experimenting with buffered channels in Go!")
// }