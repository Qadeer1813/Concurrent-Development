// --------------------------------------------
// Author: Qadeer Hussain (C00270632@setu.ie)
// Created on 14/10/2024
// Modified by: Qadeer Hussain
// --------------------------------------------
package main

import (
	"fmt"
	"sync"
	"time"
)

const size = 10

var buffer = make(chan int, size)
var wgProd = sync.WaitGroup{} // wait group for prod
var wgCons = sync.WaitGroup{} // wait group for con

func producer(id int) {
	defer wgProd.Done()
	for i := 0; i < 5; i++ { // Each producer produces 5 items.
		event := i * id
		buffer <- event // Send item to buffer.
		fmt.Printf("Producer %d produced: %d\n", id, event)
		time.Sleep(time.Millisecond * 100) // time to produce an item.
	}
}

func consumer(id int) {
	defer wgCons.Done()
	for event := range buffer { // Continues to receive until the buffer is closed.
		fmt.Printf("Consumer %d consumed: %d\n", id, event)
		time.Sleep(time.Millisecond * 150) // time to process an item.
	}
}

func main() {
	numProd := 5
	numCons := 3

	wgProd.Add(numProd)
	wgCons.Add(numCons)

	for i := 1; i <= numProd; i++ {
		go producer(i)
	}

	for i := 1; i <= numCons; i++ {
		go consumer(i)
	}

	wgProd.Wait() // Wait for all producers Finish.
	close(buffer) // Close the buffer when all producers are done.

	wgCons.Wait() // wait con finish
}
