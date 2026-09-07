package main

import (
	"fmt"
	"time"
)

type StatefulWorker struct {
	count int
	ch chan int
}

func (w *StatefulWorker) Start() {
	go func() {
		for {
			select {
			case value := <-w.ch:
				w.count += value
				fmt.Printf("Current count: %d\n", w.count)
			}
		}
	}()
}

func (w *StatefulWorker) Send(value int) {
	w.ch <- value
}

func main() {
	fmt.Println("Experimenting and understanding how stateful go routines work")
	stWorker := &StatefulWorker{
		ch: make(chan int),
	}
	stWorker.Start()
	for i := range 5 {
		stWorker.Send(i)
		time.Sleep(500 * time.Millisecond)
	}
}